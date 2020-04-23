# User Guide

## Synopsis

A command line utility tool to perform various automation tasks by APJ Build Engineering Team.

The tool was developed using [Go](https://golang.org/) with the following libraries:
- [excelize](https://github.com/360EntSecGroup-Skylar/excelize)
- [cobra](https://github.com/spf13/cobra)
- [aws-sdk-go](https://github.com/aws/aws-sdk-go/) ([official documentation](https://docs.aws.amazon.com/sdk-for-go/index.html) and [sample codes](https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code))

## Installation

```bash
go get github.com/360EntSecGroup-Skylar/excelize # Install excelize module
go get github.com/spf13/cobra                    # Install cobra module
go get github.com/spf13/cobra/cobra              # Install cobra CLI tool
go get github.com/aws/aws-sdk-go/aws             # Install aws sdk
```

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