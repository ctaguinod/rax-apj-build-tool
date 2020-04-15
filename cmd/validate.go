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
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate DD SpreadSheet",
	Long: `Validate DD Spreadsheet if all required fields are not empty. 

Example Usage:

rax-apj-build-tool validate --config config.yaml 

or 

rax-apj-build-tool validate -i ImpDoc_FAWS_APJTrial_v0.1.xlsx --sheets="Summary","Networking Services" --resources="summary","vpc","subnets"

The command will create a validated DD spreadsheet validated-ImpDoc_FAWS_APJTrial_v0.1.xlsx in current working directory.
Required cells that are empty will be highlighted in color ORANGE which means validation FAILED and needed to be filled in.
Required cells that are not empty will be highlighted in color GREEN which means validation PASS.
`,

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

		if inputFile != "" {

			// Copy inputFile to validated-inputFile and modify only validated-inputFile
			_, file := filepath.Split(inputFile)
			validatedFile := fmt.Sprintf("validated-%s", file)
			copy(inputFile, validatedFile)
			inputFile = validatedFile

			fmt.Printf("############ Validating DD Spreadsheet: %s ############\n", inputFile)
			fmt.Println()

			// Iterate each sheet
			for _, sheet := range sheets {

				for _, resource := range resources {

					if sheet == "Summary" && resource == "summary" {
						key := viper.GetString("resourcesMap.summary.key")
						values := viper.GetStringSlice("resourcesMap.summary.values")
						rows := viper.GetStringSlice("resourcesMap.summary.rows")
						if key == "" && values == nil && rows == nil {
							key := "B"
							values := []string{"C"}
							rows := []string{"10", "11", "12", "13", "16", "17", "18", "19", "20", "23", "24"}
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						} else {
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Networking Services" && resource == "vpc" {
						key := viper.GetString("resourcesMap.vpc.key")
						values := viper.GetStringSlice("resourcesMap.vpc.values")
						rows := viper.GetStringSlice("resourcesMap.vpc.rows")
						if key == "" && values == nil && rows == nil {
							key := "B"
							values := []string{"C"}
							rows := []string{"5", "6", "7", "8", "9", "10", "11", "12", "13"}
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						} else {
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Networking Services" && resource == "vpc_endpoints" {
						key := viper.GetString("resourcesMap.vpc_endpoints.key")
						values := viper.GetStringSlice("resourcesMap.vpc_endpoints.values")
						rows := viper.GetStringSlice("resourcesMap.vpc_endpoints.rows")
						if key == "" && values == nil && rows == nil {
							key := "B"
							values := []string{"C"}
							rows := []string{"21", "22", "23"}
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						} else {
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Networking Services" && resource == "subnets" {
						key := viper.GetString("resourcesMap.subnets.key")
						values := viper.GetStringSlice("resourcesMap.subnets.values")
						rows := viper.GetStringSlice("resourcesMap.subnets.rows")
						if key == "" && values == nil && rows == nil {
							key := "B"
							values := []string{"C", "D", "E", "F"}
							rows := []string{"15", "16", "17", "18", "19"}
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						} else {
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Storage & Compute Services" && resource == "ec2_instances" {
						key := viper.GetString("resourcesMap.ec2_instances.key")
						values := viper.GetStringSlice("resourcesMap.ec2_instances.values")
						rows := viper.GetStringSlice("resourcesMap.ec2_instances.rows")
						if key == "" && values == nil && rows == nil {
							key := "B"
							values := []string{"C", "D"}
							rows := []string{"17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30"}
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						} else {
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Storage & Compute Services" && resource == "auto_scaling_groups" {
						key := viper.GetString("resourcesMap.auto_scaling_groups.key")
						values := viper.GetStringSlice("resourcesMap.auto_scaling_groups.values")
						rows := viper.GetStringSlice("resourcesMap.auto_scaling_groups.rows")
						if key == "" && values == nil && rows == nil {
							key := "B"
							values := []string{"C", "D"}
							rows := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						} else {
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}
					} else if sheet == "Database" && resource == "rds" {
						key := viper.GetString("resourcesMap.rds.key")
						values := viper.GetStringSlice("resourcesMap.rds.values")
						rows := viper.GetStringSlice("resourcesMap.rds.rows")
						if key == "" && values == nil && rows == nil {
							key := "B"
							values := []string{"C"}
							rows := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						} else {
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}
					} else if sheet == "Database" && resource == "elasticache" {
						key := viper.GetString("resourcesMap.elasticache.key")
						values := viper.GetStringSlice("resourcesMap.elasticache.values")
						rows := viper.GetStringSlice("resourcesMap.elasticache.rows")
						if key == "" && values == nil && rows == nil {
							key := "B"
							values := []string{"C"}
							rows := []string{"17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34"}
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						} else {
							fmt.Printf("############ Sheet: %s ############\n", sheet)
							fmt.Printf("############ Resource: %s ############\n", resource)
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}
					}
				}

			}

		} else {
			fmt.Println("Usage: rax-apj-build-tool validate -i ImpDoc_FAWS_APJTrial_v0.1.xlsx")
		}

	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// -i flag
	validateCmd.Flags().StringP("input", "i", "", "DD Spreadsheet file to process")

	// --sheets flag
	validateCmd.Flags().StringSlice("sheets", []string{}, "Sheets to process, e.g. Networking Service, Storage & Compute Service")

	// --resources flag
	validateCmd.Flags().StringSlice("resources", []string{}, "Resources to process, e.g. vpc, subnets")
}

func validateCellsIfNotEmpty(inputFile string, sheet string, key string, columns []string, rows []string) {

	// Input File.
	xlsxFileIn, err := excelize.OpenFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Style: fill cell color to green.
	styleCellColorGreen, err := xlsxFileIn.NewStyle(`{"fill":{"type":"pattern","color":["#50C878"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}

	// Style: fill cell color to orange.
	styleCellColorOrange, err := xlsxFileIn.NewStyle(`{"fill":{"type":"pattern","color":["FFA500"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("###### Columns: %s ######\n", columns)
	fmt.Printf("###### Rows: %s ######\n", rows)
	fmt.Println()

	for _, column := range columns {

		for _, rows := range rows {
			value := xlsxFileIn.GetCellValue(sheet, fmt.Sprintf("%s%s", column, rows))
			value = strings.Replace(value, "\n", ", ", -1) //  Replace new lines with comma
			key := xlsxFileIn.GetCellValue(sheet, fmt.Sprintf("%s%s", key, rows))

			if value != "" {
				fmt.Printf("%s%s %s: %s\n", column, rows, key, value)
				xlsxFileIn.SetCellStyle(sheet, fmt.Sprintf("%s%s", column, rows), fmt.Sprintf("%s%s", column, rows), styleCellColorGreen)
			} else {
				fmt.Printf("%s%s %s: %s\n", column, rows, key, "<NULL> ***FAIL***")
				xlsxFileIn.SetCellStyle(sheet, fmt.Sprintf("%s%s", column, rows), fmt.Sprintf("%s%s", column, rows), styleCellColorOrange)
			}
		}

		fmt.Println()

	}

	// Save.
	if err = xlsxFileIn.Save(); err != nil {
		fmt.Println(err)
	}

}

// Function to copy file
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
