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

	"github.com/spf13/cobra"
)

// qcCmd represents the qc command
var qcCmd = &cobra.Command{
	Use:   "qc",
	Short: "QC AWS Build against Validated DD SpreadSheet",
	Long: `QC AWS Build against Validated DD SpreadSheet.

Example Usage:

rax-apj-build-tool qc -i validated-ImpDoc_FAWS_APJTrial_v0.1.xlsx --sheets="Summary","Networking Services"`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("qc called")
	},
}

func init() {
	rootCmd.AddCommand(qcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Look for ImpDoc_FAWS_APJTrial_v0.1.xlsx in current directory if -i flag not defined.
	qcCmd.Flags().StringP("input", "i", "validated-ImpDoc_FAWS_APJTrial_v0.1.xlsx", "Validated DD Spreadsheet file to use for QC")

	// Process sheets: "Summary", "Networking Services", "Storage & Compute Services" by default if --sheets flat not invoked.
	qcCmd.Flags().StringSlice("sheets", []string{"Summary", "Networking Services", "Storage & Compute Services"}, "Sheets to process")
}
