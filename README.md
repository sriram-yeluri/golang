# golang

## How to build go packages for multiple platforms

```go
go tool dist list

GOOS=windows go build
GOOS=linux go build

GOOS=linux GOARCH=arm64 go build
```

