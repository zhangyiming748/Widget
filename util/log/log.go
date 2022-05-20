package log

import (
	"Widget/util/conf"
	"io"
	"log"
	"os"
	"strings"
)

var (
	Info  *log.Logger
	CMD   *log.Logger
	Debug *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func init() {
	prefix := strings.Join([]string{conf.GetVal("main", "function"), ":"}, "")
	clog, err := os.OpenFile("cmd.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	wlog, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	elog, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	dlog, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	Info = log.New(os.Stdout, prefix, log.Ltime|log.Lshortfile)
	Debug = log.New(io.MultiWriter(dlog, os.Stdout), prefix, log.Ltime|log.Lshortfile)
	Warn = log.New(io.MultiWriter(wlog, os.Stdout), prefix, log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(elog, os.Stdout), prefix, log.Ltime|log.Lshortfile)
	CMD = log.New(io.MultiWriter(clog, os.Stdout), prefix, log.Ltime|log.Lshortfile)
}
