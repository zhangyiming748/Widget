package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"
	"youtube-dl/downloadcmd"
	"youtube-dl/mylog"
	"youtube-dl/readline"
	"youtube-dl/timeNow"
	"youtube-dl/util"
)

func main() {
	var (
		fp     string
		addr   string
		port   string
		target string
	)
	if ok := isWindows(); ok {
		panic(errors.New("不兼容的操作系统"))
	}
	if fp = util.GetVal("links", "path"); fp == "" {
		panic(errors.New("没有找到待下载文件列表"))
	}
	if addr = util.GetVal("proxy", "address"); addr == "" {
		panic(errors.New("没有有效的IP地址"))
	}
	if port = util.GetVal("proxy", "port"); port == "" {
		panic(errors.New("没有有效的端口"))
	}
	if target = util.GetVal("target", "fp"); target == "" {
		panic(errors.New("没有有效目标文件夹"))
	}
	proxy := strings.Join([]string{addr, port}, ":")
	fmt.Println(proxy)
	tn := timeNow.DateNowFormatStr()
	ti := time.Now()
	mylog.Logof(tn)
	mylog.Logof("\n")
	var wg sync.WaitGroup
	links := readline.Readlink(fp)
	for i, v := range links {
		wg.Add(1)
		log.Printf("开始尝试下载NO.%d\n", i)
		go downloadcmd.RunCmd(v, &wg, proxy, target,i)
	}
	wg.Wait()
	ta := timeNow.DateNowFormatStr()
	tj := time.Now()
	mylog.Logof(ta)
	mylog.Logof("\n")
	sub := tj.Sub(ti)
	log.Printf("下载完成!\t用时%v\n", sub)
}
func isWindows() bool {
	arch := runtime.GOARCH
	goos := runtime.GOOS
	if arch == "amd64" && goos == "windows" {
		//Windows
		log.Fatal("不能在Windows系统中运行,可以尝试在Windows中开启WSL")
		return true
	}
	return false
}
