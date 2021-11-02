# golang

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

