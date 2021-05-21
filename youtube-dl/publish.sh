#!/bin/bash
find ./ -name ".DS_Store" -depth -exec rm {} \;
go build -o forDarwin main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o forUbuntu main.go
CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o forWin32.exe main.go
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o forWin64.exe main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o forRaspi main.go