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
		return "Raspbian"
	}
	if arch == "amd64" || goos == "darwin" {
		//苹果
		return "macOS"
	}
	if arch == "amd64" || goos == "windows" {
		//Windows
		return "windows"
	}
	if arch == "amd64" || goos == "linux" {
		//ubuntu
		return "ubuntu"
	}
	return ""
}