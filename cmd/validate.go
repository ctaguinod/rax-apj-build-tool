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
	"reflect"
	"regexp"
	"strconv"
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
						/*
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
						*/
						// header
						fmt.Printf("############ Sheet: %s ############\n", sheet)
						fmt.Printf("############ Resource: %s ############\n", resource)
						// scan for keys
						keySlice, _ := ScanKeys(inputFile, sheet, "EC2 Standalone Instances")
						//fmt.Println("keySlice:", keySlice)

						// loop for all matched keys
						for _, v := range keySlice {
							// scan for borders
							colSlice, rowSlice := ScanBorders(inputFile, sheet, "multi", v, false)
							//fmt.Println("Cell Reference:", v, "rowSlice:", rowSlice, "with type of ", reflect.TypeOf(rowSlice))
							//fmt.Println("Cell Reference:", v, "colSlice:", colSlice, "with type of ", reflect.TypeOf(colSlice))

							// test integration
							key := colSlice[0]
							values := colSlice[1:]
							rows := rowSlice
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

/* #################################################### */
/* added by keivinc */

func ScanKeys(excelFile string, SheetTabName string, inputString string) (keySlice, regexMatchSlice []string) {

	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		println(err.Error())
		return
	}
	// search for keyword
	SearchString := "(?i)" + inputString

	// creates a empty array
	keySlice = make([]string, 0)

	// creates a regexp from string
	r, _ := regexp.Compile(SearchString)
	// regex
	//rows, err := f.GetRows(SheetTabName)
	rowsRange := f.GetRows(SheetTabName)

	// row counter
	RowCount := 0
	for _, row := range rowsRange {
		RowCount++
		// column counter
		ColCount := 0
		for _, colCell := range row {
			ColCount++

			// matches for environment
			MatchEnv := r.MatchString(colCell)

			if MatchEnv == true {
				// converts matrix to alpha-numeric cell names
				cell, err := CoordinatesToCellName(ColCount, RowCount)
				if err != nil {
					println(err.Error())
					return
				}
				// append cell coordinates to slice
				keySlice = append(keySlice, cell)

			}
		}
		//println()
	}
	fmt.Println("Matched cell values:", keySlice)

	// create slice for matched columns
	regexMatchSlice = make([]string, 0)
	// search for columns
	searchAlphabet := "[a-zA-Z]+"
	// creates a regexp from string
	alphabetRegex, _ := regexp.Compile(searchAlphabet)
	for i, v := range keySlice {
		fmt.Println("Index : value are[", i, ":", v, "]")
		regexMatches := alphabetRegex.FindStringSubmatch(v)
		fmt.Println("matched:", regexMatches)
		fmt.Println("final regexMatches:", regexMatches, "with type of", reflect.TypeOf(regexMatches))

		// append
		regexMatchSlice = append(regexMatchSlice, regexMatches[0])

		/*
			colSlice, rowSlice = ScanBorders(excelFile, SheetTabName, "multi", v, false)
			fmt.Println("Cell Reference:", v, "rowSlice:", rowSlice, "with type of ", reflect.TypeOf(rowSlice))
			fmt.Println("Cell Reference:", v, "colSlice:", colSlice, "with type of ", reflect.TypeOf(colSlice))
		*/
	}

	return

}

// ScanBorders checks if next row is empty , rowSlice []string
func ScanBorders(excelFile, sheetName, colCondition string, currentCellName string, ignoreNotes bool) (colSlice []string, rowSlice []string) {
	// Gets next RowNum
	//nextRowNum = currentRowNum + 1
	// open excel file
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		println(err.Error())
		return
	}

	/* ###################################################### */

	/* ###################################################### */

	/* breakdown cell coordinates into rows and columns */
	// search for alphabets and numbers
	searchAlphabet := "[a-zA-Z]+"
	searchNumber := "[^a-zA-Z]+"
	// creates a regexp from string
	alphabetRegex, _ := regexp.Compile(searchAlphabet)
	numberRegex, _ := regexp.Compile(searchNumber)
	// get matching sections of string
	colValue := alphabetRegex.FindStringSubmatch(currentCellName)
	rowValue := numberRegex.FindStringSubmatch(currentCellName)

	// converts []string to string and remove []
	colValueString := ConvertSliceStrToStr(colValue)
	rowValueString := ConvertSliceStrToStr(rowValue)

	// converts string to int for rows
	rowValueInt, _ := strconv.Atoi(rowValueString)
	// Checks for "row"

	// ###################################
	if colCondition == "multi" {
		endOfColBlock := false
		// slice to store range of and cols
		colSlice = make([]string, 0)
		// add initial col
		colSlice = append(colSlice, colValueString)

		// Converts Column to number
		colNum, err := ColumnNameToNumber(colValueString)
		if err != nil {
			println(err.Error())
			return
		}
		nextColValueInt := colNum + 1
		for endOfColBlock == false {
			// Converts Column to number

			nextColName, err := ColumnNumberToName(nextColValueInt)
			if err != nil {
				println(err.Error())
				return
			}
			//fmt.Println("current col is", colValueString, "with type", reflect.TypeOf(nextColName))
			//fmt.Println("next col is", nextColName, "with type", reflect.TypeOf(nextColName))

			nextCellString := nextColName + strconv.Itoa(rowValueInt)
			// get cell value
			nextCellValue := f.GetCellValue(sheetName, nextCellString)

			// append col
			colSlice = append(colSlice, nextColName)

			if nextCellValue != "" {
				nextColValueInt++
				// [DEBUG] Prints current loop
				fmt.Println("# Message:[", nextCellString, "] with value of [", nextCellValue, "]")
				continue

			} else {
				// [DEBUG] Prints current loop
				fmt.Println("# Message:[", nextCellString, "] with value [", nextCellValue, "]", "\n## End of Range ##")
				// remove the empty cell
				colSlice = colSlice[:len(colSlice)-1]

				break
			}
			return

		}

	}
	// ###################################
	// slice to store range of rows
	rowSlice = make([]string, 0)
	// add initial row
	rowSlice = append(rowSlice, strconv.Itoa(rowValueInt))
	// initial value
	nextRowValueInt := rowValueInt + 1

	endOfRowBlock := false
	for endOfRowBlock == false {
		nextCellString := colValueString + strconv.Itoa(nextRowValueInt)
		// get cell value
		nextCellValue := f.GetCellValue(sheetName, nextCellString)

		/* checks for "Notes" */
		searchNotes := "(?i)" + "^notes?"
		// creates a regexp from string
		notesRegex, _ := regexp.Compile(searchNotes)
		// get matching sections of string
		notesMatchBool := notesRegex.MatchString(nextCellValue)

		switch ignoreNotes {
		case true:
			switch notesMatchBool {
			case false:
				// append slice
				rowSlice = append(rowSlice, strconv.Itoa(nextRowValueInt))

				fmt.Println("## continue searching for notes")
			default:
				// skip row
				fmt.Println(nextCellString, "is matching Notes. Hence skipping")
			}
		default:
			rowSlice = append(rowSlice, strconv.Itoa(nextRowValueInt))
		}

		// [DEBUG] Prints current loop
		//fmt.Println("loop for:", nextCellString)
		if nextCellValue != "" {
			nextRowValueInt++
			// [DEBUG] Prints current loop
			//fmt.Println("# Message:[", nextCellString, "] with value of [", nextCellValue, "]")
			continue

		} else {
			// [DEBUG] Prints current loop
			//fmt.Println("# Message:[", nextCellString, "] with value [", nextCellValue, "]", "\n## End of Range ##")
			// remove the empty cell
			rowSlice = rowSlice[:len(rowSlice)-1]

			break
		}
		return
	} //here

	return

}

// ConvertSliceStrToStr converts []string to string and removes "[" "]"
func ConvertSliceStrToStr(inputSliceString []string) (outputString string) {
	// converts []string to string
	outputString = fmt.Sprint(inputSliceString)
	// remove []
	outputString = strings.ReplaceAll(outputString, "[", "")
	outputString = strings.ReplaceAll(outputString, "]", "")
	return
}

/* #################################################### */
/* extracted from excelize code */
func ColumnNameToNumber(name string) (int, error) {
	if len(name) == 0 {
		return -1, newInvalidColumnNameError(name)
	}
	col := 0
	multi := 1
	for i := len(name) - 1; i >= 0; i-- {
		r := name[i]
		if r >= 'A' && r <= 'Z' {
			col += int(r-'A'+1) * multi
		} else if r >= 'a' && r <= 'z' {
			col += int(r-'a'+1) * multi
		} else {
			return -1, newInvalidColumnNameError(name)
		}
		multi *= 26
	}
	return col, nil
}

// ColumnNumberToName provides a function to convert the integer to Excel
// sheet column title.
//
// Example:
//
//     excelize.ColumnNumberToName(37) // returns "AK", nil
//
func ColumnNumberToName(num int) (string, error) {
	if num < 1 {
		return "", fmt.Errorf("incorrect column number %d", num)
	}
	var col string
	for num > 0 {
		col = string((num-1)%26+65) + col
		num = (num - 1) / 26
	}
	return col, nil
}

// CellNameToCoordinates converts alphanumeric cell name to [X, Y] coordinates
// or returns an error.
//
// Example:
//
//    CellCoordinates("A1") // returns 1, 1, nil
//    CellCoordinates("Z3") // returns 26, 3, nil
//
func CellNameToCoordinates(cell string) (int, int, error) {
	const msg = "cannot convert cell %q to coordinates: %v"

	colname, row, err := SplitCellName(cell)
	if err != nil {
		return -1, -1, fmt.Errorf(msg, cell, err)
	}

	col, err := ColumnNameToNumber(colname)
	if err != nil {
		return -1, -1, fmt.Errorf(msg, cell, err)
	}

	return col, row, nil
}

// CoordinatesToCellName converts [X, Y] coordinates to alpha-numeric cell
// name or returns an error.
//
// Example:
//
//    CoordinatesToCellName(1, 1) // returns "A1", nil
//
func CoordinatesToCellName(col, row int) (string, error) {
	if col < 1 || row < 1 {
		return "", fmt.Errorf("invalid cell coordinates [%d, %d]", col, row)
	}
	colname, err := ColumnNumberToName(col)
	if err != nil {
		// Error should never happens here.
		return "", fmt.Errorf("invalid cell coordinates [%d, %d]: %v", col, row, err)
	}
	return fmt.Sprintf("%s%d", colname, row), nil
}

// SplitCellName splits cell name to column name and row number.
//
// Example:
//
//     excelize.SplitCellName("AK74") // return "AK", 74, nil
//
func SplitCellName(cell string) (string, int, error) {
	alpha := func(r rune) bool {
		return ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z')
	}

	if strings.IndexFunc(cell, alpha) == 0 {
		i := strings.LastIndexFunc(cell, alpha)
		if i >= 0 && i < len(cell)-1 {
			col, rowstr := cell[:i+1], cell[i+1:]
			if row, err := strconv.Atoi(rowstr); err == nil && row > 0 {
				return col, row, nil
			}
		}
	}
	return "", -1, newInvalidCellNameError(cell)
}
func newInvalidColumnNameError(col string) error {
	return fmt.Errorf("invalid column name %q", col)
}
func newInvalidCellNameError(cell string) error {
	return fmt.Errorf("invalid cell name %q", cell)
}
