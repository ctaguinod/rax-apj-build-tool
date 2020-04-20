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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

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
	//fmt.Println("Matched cell values:", keySlice)

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
