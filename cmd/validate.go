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

rax-apj-build-tool validate -i ImpDoc_FAWS_APJTrial_v0.1.xlsx --sheets="Summary","Networking Services"

or 

rax-apj-build-tool validate --config config.yaml 

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

				// Function validate syntax:
				// validate(inputFile, sheet, key string, columns []string, rows []int)
				if sheet == "Summary" {

					// Summary | Rows 9 to 24 | Column B C
					validateCellsIfNotEmpty(inputFile, sheet, "B", []string{"C"}, []int{10, 11, 12, 13, 16, 17, 18, 19, 20, 23, 24})

				} else if sheet == "Networking Services" {

					// Networking Services | Networking | Rows 5 to 11 | Column B C
					validateCellsIfNotEmpty(inputFile, sheet, "B", []string{"C"}, []int{5, 6, 7, 8, 9, 10, 11})

					// Networking Services | Subnets | Rows 15 to 19 | Column B C D E F
					validateCellsIfNotEmpty(inputFile, sheet, "B", []string{"C", "D", "E", "F"}, []int{15, 16, 18, 19})

				} else if sheet == "Storage & Compute Services" {

					// Storage & Compute Services | EC2 Autoscaling Groups | Rows 2 to 14 | Column B C D
					validateCellsIfNotEmpty(inputFile, sheet, "B", []string{"C", "D"}, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14})

					// Storage & Compute Services | EC2 Standalone Instances | Rows 17 to 29 | Column B C
					validateCellsIfNotEmpty(inputFile, sheet, "B", []string{"C"}, []int{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29})

				} else {
					fmt.Printf("Unknown Sheet: %s\n", sheet)
				}

			}

		} else {
			fmt.Println("Usage: rax-apj-build-tool validate -i ImpDoc_FAWS_APJTrial_v0.1.xlsx")
		}

	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Look for ImpDoc_FAWS_APJTrial_v0.1.xlsx in current directory if -i flag not defined.
	//validateCmd.Flags().StringP("input", "i", "ImpDoc_FAWS_APJTrial_v0.1.xlsx", "DD Spreadsheet file to process")
	// No Default Value
	validateCmd.Flags().StringP("input", "i", "", "DD Spreadsheet file to process")

	// Process sheets: "Summary", "Networking Services", "Storage & Compute Services" by default if --sheets flat not invoked.
	//validateCmd.Flags().StringSlice("sheets", []string{"Summary", "Networking Services", "Storage & Compute Services"}, "Sheets to process")
	// No Default Value
	validateCmd.Flags().StringSlice("sheets", []string{}, "Sheets to process")
}

func validateCellsIfNotEmpty(inputFile, sheet, key string, columns []string, rows []int) {

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

	fmt.Printf("############ Sheet: %s ############\n", sheet)
	fmt.Printf("###### Columns: %s ######\n", columns)
	fmt.Printf("###### Rows: %d ######\n", rows)
	fmt.Println()

	for _, column := range columns {

		for _, rows := range rows {
			value := xlsxFileIn.GetCellValue(sheet, fmt.Sprintf("%s%d", column, rows))
			value = strings.Replace(value, "\n", ", ", -1) //  Replace new lines with comma
			key := xlsxFileIn.GetCellValue(sheet, fmt.Sprintf("%s%d", key, rows))

			if value != "" {
				fmt.Printf("%s%d %s: %s\n", column, rows, key, value)
				xlsxFileIn.SetCellStyle(sheet, fmt.Sprintf("%s%d", column, rows), fmt.Sprintf("%s%d", column, rows), styleCellColorGreen)
			} else {
				fmt.Printf("%s%d %s: %s\n", column, rows, key, "<NULL> ***FAIL***")
				xlsxFileIn.SetCellStyle(sheet, fmt.Sprintf("%s%d", column, rows), fmt.Sprintf("%s%d", column, rows), styleCellColorOrange)
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
