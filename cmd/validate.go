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

					} else if sheet == "Networking Services" && resource == "Networking" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "Networking")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Networking Services" && resource == "Subnets" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "Subnets")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Networking Services" && resource == "VPC Endpoints" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "VPC Endpoints")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Networking Services" && resource == "VPN Gateway" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "VPN Gateway")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Networking Services" && resource == "Route53" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "Route53")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Networking Services" && resource == "Certificate Manager" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "Certificate Manager")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Storage & Compute Services" && resource == "EC2 Standalone Instances" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "EC2 Standalone Instances")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Storage & Compute Services" && resource == "EC2 Autoscaling Groups" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "EC2 Autoscaling Groups")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Storage & Compute Services" && resource == "Network Load Balancers" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "Network Load Balancers")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Storage & Compute Services" && resource == "Application Load Balancers" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "Application Load Balancers")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Storage & Compute Services" && resource == "Target Groups" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "Target Groups")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Storage & Compute Services" && resource == "S3 Buckets" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "S3 Buckets")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Storage & Compute Services" && resource == "EFS" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "EFS")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Database" && resource == "RDS" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "RDS")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}

					} else if sheet == "Database" && resource == "Elasticache - Redis" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "Elasticache - Redis")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}
					} else if sheet == "Security Groups" && resource == "Security Group" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "Security Group")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}
					} else if sheet == "IAM" && resource == "IAM Resource" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "IAM Resource")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
							validateCellsIfNotEmpty(inputFile, sheet, key, values, rows)
						}
					} else if sheet == "Additional Services" && resource == "CodeDeploy" {
						// Header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)

						// Scan for Keys
						keySlice, _ := ScanKeys(inputFile, sheet, "CodeDeploy")

						// Loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
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
