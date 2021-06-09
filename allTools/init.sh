#!/bin/bash
find . -type f -name "*.log" -exec rm {} \;
find . -name "*.go" -exec gofmt -w {} \;
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o forLinux main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o forRaspi main.go
CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o forWin32.exe main.go
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o forWin64.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o forMac main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o forM1 main.go
