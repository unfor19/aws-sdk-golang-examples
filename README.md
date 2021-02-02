# aws-sdk-golang-examples

An example of how to use AWS SDK Resource Groups Tagging API (resourcegroupstaggingapi) in GoLang.

Tested on [WSL2](https://docs.microsoft.com/en-us/windows/wsl/install-win10) Ubuntu 20.04

## Requirements

- [golang v1.15](https://golang.org/doc/install)
- [AWS Credentials](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html)
- (Optional) [Visual Studio Code](https://code.visualstudio.com/download) and its [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Installation

1. Prepare your terminal's ~/.rc file (`.bashrc`, `.bash_profile`, `.zshrc`, etc.)

    ```bash
    # GOLANG
    export GOPATH="$HOME/go"
    export GOBIN="/usr/local/go/bin"
    export PATH="$PATH:$GOPATH:$GOBIN"
    # export GO111MODULE=on # IMPORTANT: If you see this line, comment it out
    ```

1. Reload the rc file - `source ~/.bashrc`
1. Just in case - `unset GO111MODULE`
1. Create a directory in `${HOME}/go/`
    ```bash
    $ TARGET_DIR="${HOME}/go/src/internal/aws-sdk-golang-resourcegroupstaggingapi"
    $ mkdir -p "$TARGET_DIR"
    ```
1. Why `/internal/`? Using the reserved keyword `internal` as a directory name instructs `go get` to consider all sub-directories as [internal packages](https://golang.org/doc/go1.4#internalpackages)

## Usage

```bash
# Download and install dependencies
$ TARGET_DIR="${HOME}/go/src/internal/aws-sdk-golang-resourcegroupstaggingapi"
$ cd "$TARGET_DIR"
$ cp aws-sdk-golang-resourcegroupstaggingapi.go "$TARGET_DIR"/
$ go get -u -v # -u = download latest, -v = verbose output

# `go get` Also builds a binary and copies it to /usr/local/go/bin/aws-sdk-golang-examples
# TODO: figure out why

# Edit aws-sdk-golang-resourcegroupstaggingapi.go
# Set the relevant region and filters

# Run the application from source code
$ go run .

# Build the application and run the binary
$ go build .
$ ./aws-sdk-golang-examples

# Check the size of the binary file
$ du -h aws-sdk-golang-examples # du = disk usage, -h = human readable
8.8M
```


## Troubleshooting

1. `go get` permission denied to `/usr/local/go/bin` directory

    **Error**
    ```bash
    go get internal/aws-sdk-golang-examples: copying /tmp/go-build594918589/b001/exe/a.out: open /usr/local/go/bin/aws-sdk-golang-examples: permission denied
    ```

    **Fix**
    ```bash
    # Set current user:user_group as the owner
    $ sudo chown -R $USER:$(id -g -n $USER) /usr/local/go/bin/
    ```