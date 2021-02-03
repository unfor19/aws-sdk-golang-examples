# aws-sdk-golang-examples

An example of how to use AWS SDK Resource Groups Tagging API (resourcegroupstaggingapi) in GoLang.

Especially useful if you're new go GoLang (like me). Tested on [WSL2](https://docs.microsoft.com/en-us/windows/wsl/install-win10) Ubuntu 20.04

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
# You can also use `go mod download`
# See `go.mod` and `go.sum` section and check the Dockerfile

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

## `go.mod` and `go.sum`

To generate [go.mod](./go.mod) and [go.sum](./go.sum), follow the next steps

```bash
# Adds current package to `go.mod`
$ go mod init

# Adds dependencies and creates `go.sum` (lock file)
$ go mod tidy

# Download dependencies according to go.mod
$ go mod download
```

## Docker

```bash
# Build
$ docker build -t unfor19/aws-sdk-golang-examples:resourcegroupstaggingapi .

# Generate `.env` file
$ printenv | grep AWS_ > .env

# Run - Passing credentials with `--env-file`
docker run --rm -it --env-file=.env unfor19/aws-sdk-golang-examples:resourcegroupstaggingapi

# Check Docker image size
docker system df -v | grep "REPOSITORY\|resourcegroupstaggingapi"
REPOSITORY                        TAG                        IMAGE ID       CREATED             SIZE      SHARED SIZE   UNIQUE SIZE   CONTAINERS
unfor19/aws-sdk-golang-examples   resourcegroupstaggingapi   cb130f3efcfe   7 minutes ago       23.88MB   14.74MB       9.136MB       0

#      real size:  23.88MB
#    shared size:  14.75MB
#    unique size:   9.14MB
# DockerHub size:  11.67MB (https://hub.docker.com/r/unfor19/aws-sdk-golang-examples/tags?page=1&ordering=last_updated)
```

## Troubleshooting

1. `go get` permission denied to `/usr/local/go/bin` directory

    **Error**
    ```bash
    go get internal/aws-sdk-golang-examples: copying /tmp/go-build594918589/b001/exe/a.out: open /usr/local/go/bin/aws-sdk-golang-examples: permission denied

    # Human friendly
    go get internal/aws-sdk-golang-examples: 
    copying /tmp/go-build594918589/b001/exe/a.out: 
    open /usr/local/go/bin/aws-sdk-golang-examples: permission denied
    ```

    **Fix**
    ```bash
    # Set current user:user_group as the owner
    $ sudo chown -R $USER:$(id -g -n $USER) /usr/local/go/bin/
    ```
1. App fails to run due to credentials issue

    **Error**
    ```bash
    failed to list resources, operation error Resource Groups Tagging API: GetResources, exceeded maximum number of attempts, 3, https response error StatusCode: 0, RequestID: , request send failed, Post "https://tagging.eu-west-1.amazonaws.com/": x509: certificate signed by unknown authority

    # Human friendly
    failed to list resources,
    operation error Resource Groups Tagging API: GetResources, 
    exceeded maximum number of attempts, 3, 
    https response error StatusCode: 0,
    RequestID: , request send failed,
    Post "https://tagging.eu-west-1.amazonaws.com/":
    x509: certificate signed by unknown authority
    ```

    **Fix** - Using [alpine](https://hub.docker.com/_/alpine/) for the final image, and not [scratch](https://hub.docker.com/_/scratch/) because `AWS_SESSION_TOKEN` env var doesn't work well on scratch.
