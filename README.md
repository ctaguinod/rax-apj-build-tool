# rax-apj-build-tool

## What is rax-apj-build-tool

A command line utility tool to perform various automation tasks by APJ Build Engineering Team.

``` bash
./rax-apj-build-tool validate --config config.yaml # Validate DD Spreadsheet if all required fields are not empty
./rax-apj-build-tool build --config config.yaml    # Generate build parameters (.tf. .tfvars) from validated DD spreadsheet
./rax-apj-build-tool qc --config config.yaml       # Automated QC of AWS Environment using validated DD spreadsheet
```

Read the [User Guide](./docs/USER_GUIDE.md) and [Developer Guide](./docs/DEVELOPER_GUIDE.md) for detailed documentation.

## Installation

Binaries are available on the [releases](https://github.com/ctaguinod/rax-apj-build-tool/releases) page. To install, download the binary for your platform from "Assets" and place this into your `$PATH`:

```bash
curl -Lo ./terraform-docs https://github.com/segmentio/terraform-docs/releases/download/v0.9.1/terraform-docs-v0.9.1-$(uname | tr '[:upper:]' '[:lower:]')-amd64
chmod +x ./terraform-docs
mv ./rax-apj-build-tool /some/dir/to/your/PATH/rax-apj-build-tool
```

**NOTE:** Windows releases are in `EXE` format.