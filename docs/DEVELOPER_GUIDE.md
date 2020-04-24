# Developer Guide

## Synopsis

A command line utility tool to perform various automation tasks by APJ Build Engineering Team.

The tool was developed using [Go](https://golang.org/) with the following libraries:
- [excelize](https://github.com/360EntSecGroup-Skylar/excelize) - Golang library for reading and writing Microsoft Excel™ (XLSX) files.
- [cobra](https://github.com/spf13/cobra) - Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.
- [aws-sdk-go](https://github.com/aws/aws-sdk-go/) - Official AWS SDK for the Go programming language. ([official documentation](https://docs.aws.amazon.com/sdk-for-go/index.html) and [sample codes](https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code)).

## Module Installation

```bash
go get github.com/360EntSecGroup-Skylar/excelize # Install excelize module
go get github.com/spf13/cobra                    # Install cobra module
go get github.com/spf13/cobra/cobra              # Install cobra CLI tool
go get github.com/aws/aws-sdk-go/aws             # Install aws sdk
```

## Commands

### Validate

The [`./rax-apj-build-tool validate`](../cmd/validate.go) command Validates DD Spreadsheet if all required fields are not empty.

[`validate`](../cmd/validate.go) uses [excelize](https://github.com/360EntSecGroup-Skylar/excelize) to Read Input DD Spreadsheet and Output Validated Spreadsheet.

Functions:

- `copy` - Used to copy input file spreadsheet to validated spreadsheet.
- `ScanKeys` and `ScanBorders` - Used to scan the input spreadsheet file, sheet and search for the resource name using regex and will output the cell ranges to be used and processed by `validateCellsIfNotEmpty`.
- `validateCellsIfNotEmpty` - Used to validate cells if empty, non empty cells will highlight the cell to color GREEN and empty cell will highlight the color to ORANGE meaning needs to be filled in.

Flags:

- `-i` - Input File needs to be .xlsx spreadsheet.
- `--sheets` - The sheets to process.
- `--resources` - The resources (Keys) to process.

### QC

The [`./rax-apj-build-tool qc`](../cmd/qc.go) command runs an automated QC of AWS Environment using the validated DD spreadsheet.

[`qc`](../cmd/qc.go) uses [excelize](https://github.com/360EntSecGroup-Skylar/excelize) to Read Input DD Spreadsheet and Outputs result if PASS or FAIL in CLI.

Functions:

- `ScanKeys` and `ScanBorders` - Used to scan the input spreadsheet file, sheet and search for the resource name using regex and will output the cell ranges to be used and processed by `getResourcesMap`.
- `getResourcesMap` - Used to scan the cell values and outputs it nested map to be used to query the AWS environment if the resource exists.
- `getVpc` - Query the AWS environment for the VPC based on the DD Spreadsheet.
- `getSubnets` - Query the AWS environment for the Subnets based on the DD Spreadsheet.
- `getEC2` - Query the AWS environment for the EC2 instances based on the DD Spreadsheet.

Flags:

- `-i` - Input File needs to be .xlsx spreadsheet.
- `--sheets` - The sheets to process.
- `--resources` - The resources (Keys) to process.

## Cobra Basic

How to initialize Cobra on new project/folder.

```
mkdir my-cobra-project
cd my-cobra-project
cobra init --pkg-name my-cobra-project
go mod init my-cobra-project
```

How to add cobra command.

```
cobra add <commandName>
cobra add validate
cobra add qc
```

## Git Cheatsheets

```bash
git tag -l                       # List tags
git tag v0.0.1                   # Create tag v0.0.1
git push --tags                  # Push tag to origin
git tag -d v0.0.1                # Delete tag in local
git push --delete origin v0.0.1  # Delete tag in origin
git add <file or folder>         # Add file/folder to Git
git commit -m <commit message>   # Commit change
git push                         # Push to origin
git status                       # Check git status
```

## Makefile

```bash
make help                   # Show available commands
make tools                  # Install required tools
make build                  # Build binary file for current platform, output will be stored in ./bin/
make build-all              # Build binary files for linux mac and windows platform, output will be stored in ./bin/
make release version=0.1.0  # Release version, *** Not fully tested, currently not working on Mac ***
```
