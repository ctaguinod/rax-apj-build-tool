# User Guide

## Synopsis

A command line utility tool to perform various automation tasks by APJ Build Engineering Team.

## Options, Syntax, Usage, and Outputs

```
Usage:
  rax-apj-build-tool [command]

Available Commands:
  help        Help about any command
  qc          QC AWS Build against Validated DD SpreadSheet
  validate    Validate DD SpreadSheet

Flags:
      --config string   sample config files in https://github.com/ctaguinod/rax-apj-build-tool/blob/master/examples/
  -h, --help            help for rax-apj-build-tool

Use "rax-apj-build-tool [command] --help" for more information about a command.
```

## Validate

The `./rax-apj-build-tool validate` command Validates DD Spreadsheet if all required fields are not empty.

The `validate` command will generate a new spreadsheet prefixed with `validated-` in the file name, e.g. if the input file name used is `ImpDoc_FAWS.xlsx` the validated output file will be `validated-ImpDoc_FAWS.xlsx`.

The validated DD spreadsheet will be updated with cells highlighted in color `GREEN` for `PASS` and `ORANGE` for `FAILED`. PASS means required field is properly filled in, FAILED means required field is left blank and should be filled in.

Example [config.yaml](https://gist.github.com/ctaguinod/0cf0f1091ac2733435692d776cbfbb0d) and DD Spreadsheets files are provided in the [examples](https://github.com/ctaguinod/rax-apj-build-tool/blob/master/examples/) directory.

```
Example Usage:

./rax-apj-build-tool validate --config config.yaml 

or 

./rax-apj-build-tool validate -i ImpDoc_FAWS.xlsx --sheets="Networking Services","Storage & Compute Services" --resources="Networking","Subnetworks","EC2 Standalone Instances","EC2 Autoscaling Groups"

The command will create a validated DD spreadsheet validated-ImpDoc_FAWS.xlsx in current working directory.
Required cells that are empty will be highlighted in color ORANGE which means validation FAILED and needed to be filled in.
Required cells that are not empty will be highlighted in color GREEN which means validation PASS.

Usage:
  rax-apj-build-tool validate [flags]

Flags:
  -h, --help                help for validate
  -i, --input string        DD Spreadsheet file to process
      --resources strings   Resources to process, e.g. Networking, Subnetworks, EC2 Standalone Instances
      --sheets strings      Sheets to process, e.g. Networking Service, Storage & Compute Service

Global Flags:
      --config string   sample config files in https://github.com/ctaguinod/rax-apj-build-tool/blob/master/examples/
```

## QC

The `./rax-apj-build-tool qc` command runs an automated QC of AWS Environment using the validated DD spreadsheet.

The `qc` command requires valid aws credentials before running. 

The `qc` command runs a QC check for provisioned resources in the actual AWS environment againts the DD spreadsheet.

Example [config.yaml](https://gist.github.com/ctaguinod/65a39d6e3df027d626ce30878f05b9a1) and DD Spreadsheets files are provided in the [examples](https://github.com/ctaguinod/rax-apj-build-tool/blob/master/examples/) directory.


```
Example Usage:

./rax-apj-build-tool qc --config config.yaml 

or 

./rax-apj-build-tool qc -i validated-ImpDoc_FAWS.xlsx --sheets="Networking Services","Storage & Compute Services" --resources="Networking","Subnetworks","EC2 Standalone Instances"

Usage:
  rax-apj-build-tool qc [flags]

Flags:
  -h, --help                help for qc
  -i, --input string        DD Spreadsheet file to process
      --resources strings   Resources to process, e.g. Networking, Subnetworks, EC2 Standalone Instances
      --sheets strings      Sheets to process, e.g. Networking Service, Storage & Compute Service

Global Flags:
      --config string   sample config files in https://github.com/ctaguinod/rax-apj-build-tool/blob/master/examples/
```
