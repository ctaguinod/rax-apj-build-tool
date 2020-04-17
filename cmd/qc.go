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

		// Get the --resources flag value
		resources, _ := cmd.Flags().GetStringSlice("resources")
		if viper.GetStringSlice("resources") != nil {
			resources = viper.GetStringSlice("resources")
		}

		if inputFile != "" {

			fmt.Printf("############ QC for DD Spreadsheet: %s ############\n", inputFile)
			fmt.Println()

			// Iterate each sheet

			for _, resource := range resources {

				if resource == "vpc" {

					fmt.Printf("############ Resource: %s ############\n", resource)

					/*
						sheet := viper.GetString("resourcesMap.vpc.sheet")
						key := viper.GetString("resourcesMap.vpc.key")
						values := viper.GetStringSlice("resourcesMap.vpc.values")
						rows := viper.GetStringSlice("resourcesMap.vpc.rows")
						if sheet == "" && key == "" && values == nil && rows == nil {
							sheet := "Networking Services"
							key := "B"
							values := []string{"C"}
							rows := []string{"5", "6", "7", "8", "9", "10", "11", "12", "13"}
							fmt.Printf("############ Default Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							qcVpc(inputFile, sheet, key, values, rows)
						} else {
							fmt.Printf("############ Else Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							//qcVpc(inputFile, sheet, key, values, rows)
							result := qcVpc(inputFile, sheet, key, values, rows)
							fmt.Println(result)
							//fmt.Println(result["Networking"])
							//fmt.Println(result["Name of Environment"])
						}
					*/
					sheet := viper.GetString("resourcesMap.vpc.sheet")
					// Scan for Keys
					keySlice, _ := ScanKeys(inputFile, sheet, "VPC Subnets")

					// Loop for all matched keys
					for _, v := range keySlice {
						// scan for borders
						colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
						key := colSlice[0]
						values := colSlice[1:]
						rows := rowSlice

						data := qcVpc(inputFile, sheet, key, values, rows)

						//fmt.Println(colSlice)
						//fmt.Println(rowSlice)
						//fmt.Println(key)
						//fmt.Println(values)
						//fmt.Println(rows)

						fmt.Println(data)
						fmt.Println(data["C"])
						//fmt.Println(data["C"]["CIDR"])
						fmt.Println(data["D"])
						//fmt.Println(data["D"]["CIDR"])
						fmt.Println(data["E"])
						//fmt.Println(data["E"]["CIDR"])
						fmt.Println(data["F"])
						//fmt.Println(data["F"]["CIDR"])

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

func qcVpc(inputFile string, sheet string, key string, columns []string, rows []string) map[string]map[string]string {

	// Input File.
	xlsxFileIn, err := excelize.OpenFile(inputFile)
	if err != nil {
		fmt.Println(err)
		//return
	}

	var columnKey string
	var columnValue string
	mainMap := make(map[string]map[string]string)

	fmt.Printf("###### Columns: %s ######\n", columns)
	fmt.Printf("###### Rows: %s ######\n", rows)
	fmt.Println()

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

	//fmt.Println(mainMap)
	return mainMap
}
