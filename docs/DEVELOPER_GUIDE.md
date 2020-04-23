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
make help      # Show available commands
make tools     # Install required tools
make build     # Build binary file for current platform, output will be stored in ./bin/
make build-all # Build binary files for linux mac and windows platform, output will be stored in ./bin/
```