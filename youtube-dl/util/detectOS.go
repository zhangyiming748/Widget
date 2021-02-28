package util

import (

	"runtime"
)

//探测当前操作系统

func DetectOS() string{
	arch := runtime.GOARCH

	goos := runtime.GOOS

	if arch == "arm" || goos == "linux" {
		//树莓派
		return "192.168.1.7:8889"
	}
	if arch == "amd64" || goos == "darwin" {
		//苹果
		return "127.0.0.1:8889"
	}
	if arch == "amd64" || goos == "windows" {
		//Windows
		return "192.168.1.7:8889"
	}
	if arch == "amd64" || goos == "linux" {
		//ubuntu
		return "192.168.1.7:8889"
	}
	return ""
}