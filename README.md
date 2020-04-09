# rax-apj-build-tool
APJ Build CLI Tool

## Example Go commands

```
go run main.go
go fmt
go build 
```

## Modules Used

### Excelize

```
go get github.com/360EntSecGroup-Skylar/excelize
```


### Cobra 
- Cobra Syntax: APPNAME Command Args --Flags or APPNAME Command --Flags Args  (rax-apj-build-tool <command> [arguments] --flag)

```
go get github.com/spf13/cobra
go get -u github.com/spf13/cobra
```

Install Cobra CLI.

```
go get github.com/spf13/cobra/cobra
go get -u github.com/spf13/cobra/cobra
```

Example to initialize Cobra on new project/folder.

```
mkdir my-cobra-project
cd my-cobra-project
cobra init --pkg-name my-cobra-project
go mod init my-cobra-project
```

To add cobra command

```
cobra add <commandName>
cobra add validate
cobra add qc
```

### AWS SDK

```
go get github.com/aws/aws-sdk-go/aws
go get github.com/jmespath/go-jmespath
```


## CLI Tool Features / Functionalities:
- Required Input: DD Spreadsheet.
- Tool Name: rax-apj-build-tool
    - Command: validate (e.g. rax-apj-build-tool validate -i i.xlsx -s "Networking Services")
        - Flags: 
            - --input | -i: DD SpreadSheet, xlsx file
            - --sheet : Sheet name to process
            - --config : YAML config file
    - Command: qc (e.g. rax-apj-build-tool qc -i i.xlsx -s "Networking Services")
        - Flags: 
            - --input | -i: validated xlsx file
            - --sheet : Sheet name to process

## ToDos
- Currently hardcoded rows/columns range to process, need to find way to be more extendible, e.g. external config/schema in YAML file.
- For QC: use sdk-for-go to query resources from DD Spreadsheet
    - https://docs.aws.amazon.com/sdk-for-go/index.html
    - https://docs.aws.amazon.com/sdk-for-go/api/index.html
    - https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code
    - https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.DescribeVpcs
    - https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#EC2.DescribeSubnets