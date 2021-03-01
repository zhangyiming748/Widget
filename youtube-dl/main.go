package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
	"youtube-dl/downloadcmd"
	"youtube-dl/mylog"
	"youtube-dl/readline"
	"youtube-dl/timeNow"
	. "youtube-dl/util"
)

func main() {
	var (
		fp string
		addr string
		port string
		target string
	)
	if fp = GetVal("links","path"); fp == "no value" {
		panic(errors.New("没有找到待下载文件列表"))
	}
	if addr=GetVal("proxy","address");addr =="no value"{
		panic(errors.New("没有有效的IP地址"))
	}
	if port=GetVal("proxy","port");port =="no value"{
		panic(errors.New("没有有效的端口"))
	}
	if target=GetVal("target","fp");addr =="no value"{
		panic(errors.New("没有有效目标文件夹"))
	}
	proxy:=strings.Join([]string{addr,port},":")
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
		go downloadcmd.RunCmd(v, &wg, proxy, target)
	}
	wg.Wait()
	ta := timeNow.DateNowFormatStr()
	tj:=time.Now()
	mylog.Logof(ta)
	mylog.Logof("\n")
	sub:=tj.Sub(ti)
	log.Printf("下载完成!\t用时%v\n",sub)
}