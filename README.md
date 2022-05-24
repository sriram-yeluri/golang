# golang

### Set GOPATH

```
# for macOS
# vi .zshrc
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=${PATH}:$GOBIN
```

### Install cobra-cli

[cora-cli-readme](https://github.com/spf13/cobra-cli/blob/main/README.md)

```go
go install github.com/spf13/cobra-cli@latest
```

### How to format go code 

```go
# format all files in a directory (including sub directories)
go fmt ./...
```

### How to build go packages for multiple platforms

[go-docs-env-vars](https://golang.org/doc/install/source#environment)

```go
# get the list of all supported distributions
go tool dist list

GOOS=windows go build

GOOS=linux go build

GOOS=linux GOARCH=arm64 go build

GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe app.go
```

