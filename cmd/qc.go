/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// qcCmd represents the qc command
var qcCmd = &cobra.Command{
	Use:   "qc",
	Short: "QC AWS Build against Validated DD SpreadSheet",
	Long: `QC AWS Build against Validated DD SpreadSheet.

Example Usage:

rax-apj-build-tool qc -i validated-ImpDoc_FAWS_APJTrial_v0.1.xlsx --resources="summary","vpc","subnets"`,

	Run: func(cmd *cobra.Command, args []string) {

		// Get the -i flag value
		inputFile, _ := cmd.Flags().GetString("input")
		if viper.GetString("input") != "" {
			inputFile = viper.GetString("input")
		}

		// Get the --sheets flag value
		sheets, _ := cmd.Flags().GetStringSlice("sheets")
		if viper.GetStringSlice("sheets") != nil {
			sheets = viper.GetStringSlice("sheets")
		}

		// Get the --resources flag value
		resources, _ := cmd.Flags().GetStringSlice("resources")
		if viper.GetStringSlice("resources") != nil {
			resources = viper.GetStringSlice("resources")
		}

		var environment string
		var region string
		var vpcName string

		if inputFile != "" {

			fmt.Printf("############ QC for DD Spreadsheet: %s ############\n", inputFile)
			// Iterate each sheet
			for _, sheet := range sheets {

				for _, resource := range resources {

					// Scan for Keys
					keySlice, _ := ScanKeys(inputFile, sheet, resource)
					for _, v := range keySlice {

						// scan for borders
						colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
						key := colSlice[0]
						values := colSlice[1:]
						rows := rowSlice

						resourcesMap := getResourcesMap(inputFile, sheet, key, values, rows)
						//fmt.Println(resourcesMap)

						// VPC
						if resource == "Networking" {
							for _, column := range values {
								region = resourcesMap[column]["Region"]
								environment = resourcesMap[column]["Name of Environment"]
								vpcName = resourcesMap[column]["Networking"]
								vpcCidr := resourcesMap[column]["CIDR Range"]
								fmt.Println()
								fmt.Printf("############ QC for VPC ############\n")
								fmt.Printf("###### DD Parameters ######\n")
								fmt.Printf("Region: %s\n", region)
								fmt.Printf("Environment: %s\n", environment)
								fmt.Printf("VPC Name: %s\n", vpcName)
								fmt.Printf("VPC CIDR: %s\n", vpcCidr)
								fmt.Println()
								getVpc(environment, region, vpcName, vpcCidr)
							}
						}

						//Subnets
						if resource == "Subnetworks" {
							for _, column := range values {
								subnetName := resourcesMap[column]["Subnetworks"]
								subnetName = fmt.Sprintf("%s-%s", vpcName, subnetName)
								subnetCidr := resourcesMap[column]["CIDR"]
								subnetAz := resourcesMap[column]["AZ"]
								fmt.Println()
								fmt.Printf("############ QC for Subnet %s ############\n", column)
								fmt.Printf("###### DD Parameters ######\n")
								fmt.Printf("Region: %s\n", region)
								fmt.Printf("Environment: %s\n", environment)
								fmt.Printf("Subnet Name: %s\n", subnetName)
								fmt.Printf("Subnet CIDR: %s\n", subnetCidr)
								fmt.Printf("AZ: %s\n", subnetAz)
								fmt.Println()
								getSubnets(environment, region, subnetName, subnetCidr, subnetAz)
							}
						}

						// EC2 Standalone Instances
						if resource == "EC2 Standalone Instances" {
							for _, column := range values {
								instanceName := resourcesMap[column]["EC2 Standalone Instances"]
								instanceType := resourcesMap[column]["Instance Type"]
								instanceImage := resourcesMap[column]["Instance Image"]
								instanceSubnet := resourcesMap[column]["Subnet"]
								rootVolume := resourcesMap[column]["Root Volume Size (GB)"]
								rootVolumeEncrypt := resourcesMap[column]["Encrypt Root Volume"]
								dataVolume := resourcesMap[column]["Data Volume Size(s) (GB)"]
								dataVolumeEncrypt := resourcesMap[column]["Encrypt Data Volumes"]
								securityGroups := resourcesMap[column]["SecurityGroup(s)"]
								fmt.Println()
								fmt.Printf("############ QC for EC2 Standalone Instance %s ############\n", column)
								fmt.Printf("###### DD Parameters ######\n")
								fmt.Printf("Region: %s\n", region)
								fmt.Printf("Environment: %s\n", environment)
								fmt.Printf("EC2 Instance Name: %s\n", instanceName)
								fmt.Printf("Instance Image: %s\n", instanceImage)
								fmt.Printf("Subnet: %s\n", instanceSubnet)
								fmt.Printf("Root Volume Size (GB): %s\n", rootVolume)
								fmt.Printf("Encrypt Root Volume: %s\n", rootVolumeEncrypt)
								fmt.Printf("Data Volume Size(s) (GB): %s\n", dataVolume)
								fmt.Printf("Encrypt Data Volumes: %s\n", dataVolumeEncrypt)
								fmt.Printf("SecurityGroup(s): %s\n", securityGroups)
								fmt.Println()
								getEC2(environment, region, instanceName, instanceType, rootVolume, rootVolumeEncrypt)
							}
						}

					}
				}

			}

		} else {
			fmt.Println("Usage: rax-apj-build-tool qc --config config.yaml")
		}

	},
}

func init() {
	rootCmd.AddCommand(qcCmd)

	// -i flag
	qcCmd.Flags().StringP("input", "i", "", "DD Spreadsheet file to process")

	// --resources flag
	qcCmd.Flags().StringSlice("resources", []string{}, "Resources to process, e.g. vpc, subnets")
}

func getResourcesMap(inputFile string, sheet string, key string, columns []string, rows []string) map[string]map[string]string {

	// Input File.
	xlsxFileIn, err := excelize.OpenFile(inputFile)
	if err != nil {
		fmt.Println(err)
		//return
	}

	var columnKey string
	var columnValue string
	mainMap := make(map[string]map[string]string)

	//fmt.Printf("###### Columns: %s ######\n", columns)
	//fmt.Printf("###### Rows: %s ######\n", rows)
	//fmt.Println()

	for _, column := range columns {

		for _, row := range rows {
			columnValue = xlsxFileIn.GetCellValue(sheet, fmt.Sprintf("%s%s", column, row))
			columnValue = strings.Replace(columnValue, "\n", ", ", -1) //  Replace new lines with comma
			columnKey = xlsxFileIn.GetCellValue(sheet, fmt.Sprintf("%s%s", key, row))

			contentsMap, err := mainMap[column]
			if columnValue != "" {
				if !err {
					contentsMap = make(map[string]string)
					mainMap[column] = contentsMap
				}
				contentsMap[columnKey] = columnValue

			} else {
				if !err {
					contentsMap = make(map[string]string)
					mainMap[column] = contentsMap
				}
				contentsMap[columnKey] = " <NULL> "
			}

		}

	}

	return mainMap
}
