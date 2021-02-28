package main

import (
	"log"
	"os"
	"sync"
	"youtube-dl/downloadcmd"
	"youtube-dl/mylog"
	"youtube-dl/readline"
	"youtube-dl/timeNow"
	"youtube-dl/util"
)

const (
	list = "links.txt"
)

func main() {
	var fp string
	if fp=util.GetArgs();fp==""{
		fp=list
	}
	proxy:=util.DetectOS()
	path:=util.GetExcPath()
	tn := timeNow.DateNowFormatStr()
	mylog.Logof(tn)
	mylog.Logof("\n")
	var wg sync.WaitGroup
	links := readline.Readlink(fp)
	for i, v := range links {
		wg.Add(1)
		log.Printf("开始尝试下载NO.%d\n", i)
		go downloadcmd.RunCmd(v, &wg,proxy,path)
	}
	wg.Wait()
	ta := timeNow.DateNowFormatStr()
	mylog.Logof(ta)
	mylog.Logof("\n")
}
func deleteFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Remove(filename)
		return false
	}
	return true
}
