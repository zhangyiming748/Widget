package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"youtube-dl/downloadcmd"
	. "youtube-dl/mylog"
	"youtube-dl/readline"
	"youtube-dl/timeNow"
	"youtube-dl/util"
)

var (
	MaxGoroutine = util.GetVal("goroutine", "num")
)

func init() {
	if isExists("error.txt") {
		if err := os.Remove("error.txt"); err != nil {
			Error.Println("初始化错误日志失败")
		}
	}
	if isExists("debug.txt") {
		if err := os.Remove("debug.txt"); err != nil {
			Error.Println("初始化调试日志失败")
		}
	}
}
func main() {
	var (
		fp     string
		addr   string
		port   string
		target string
	)
	defer func() {
		if err := recover(); err != nil {
			Debug.Printf("程序运行过程中产生的错误:%v", err)
		}
	}()
	if ok := isWindows(); ok {
		panic(errors.New("不兼容的操作系统"))
	}
	if fp = util.GetVal("links", "path"); fp == "" {
		panic(errors.New("没有找到待下载文件列表"))
	}
	if addr = util.GetVal("proxy", "address"); addr == "" {
		panic(errors.New("没有有效的ip"))
	}
	if port = util.GetVal("proxy", "port"); port == "" {
		panic(errors.New("没有有效的端口"))
	}
	if target = util.GetVal("target", "fp"); target == "" {
		panic(errors.New("没有有效目标文件夹"))
	}
	if MaxGoroutine == "no value" || MaxGoroutine == "" {
		MaxGoroutine = "1"
		Debug.Println("没有定义最大连接数,默认为单线程")
	}
	proxy := strings.Join([]string{addr, port}, ":")
	fmt.Println(proxy)
	tn := timeNow.DateNowFormatStr()
	ti := time.Now()
	Debug.Println(tn)
	//Debug.Println("\n")
	var wg sync.WaitGroup
	max, _ := strconv.Atoi(MaxGoroutine)
	ch := make(chan struct{}, max)
	links := readline.Readlink(fp)
	list := make(map[string]bool)
	for i, v := range links {
		if list[v] == true {
			Debug.Printf("跳过重复文件No.%d", i+1)
			continue
		}
		ch <- struct{}{}
		wg.Add(1)
		Debug.Printf("开始尝试下载NO.%d\n", i+1)
		go downloadcmd.RunCmd(v, &wg, proxy, target, i, ch)
		list[v] = true
	}
	wg.Wait()
	ta := timeNow.DateNowFormatStr()
	tj := time.Now()
	Debug.Println(ta)
	sub := tj.Sub(ti)
	Debug.Printf("下载完成!\t用时%v\n", sub)
}
func isWindows() bool {
	arch := runtime.GOARCH
	goos := runtime.GOOS
	if arch == "amd64" && goos == "windows" || arch == "386" && goos == "windows" {
		//Windows
		Debug.Fatal("不能在Windows系统中运行,可以尝试在Windows中开启WSL")
		return true
	}
	return false
}
func isExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
