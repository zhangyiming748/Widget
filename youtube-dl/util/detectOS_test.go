package util

import (
	"log"
	"runtime"
	"testing"
)

func TestDetectOS(t *testing.T) {
	a:=runtime.GOARCH
	// 返回当前的系统架构
	b:=runtime.GOOS//
	// 返回当前的操作系统
	log.Println(a,b)
	//pi=arm linux
	//MacOS=amd64 darwin
	//Windows=amd64 windows
	//WSL=amd64 linux
	//ubuntu=amd64 linux
}
