#!/bin/bash
go build -o weather4darwin main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather4ubuntu main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o weather4windows main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o weather4raspi main.go
