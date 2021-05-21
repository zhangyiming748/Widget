package mylog

import (
	"io"
	"log"
	"os"
)

var (
	Info *log.Logger
	Debug *log.Logger
	Error *log.Logger
)

func init() {
	log.SetPrefix("youtube_dl: ")
	log.SetFlags(log.Ltime | log.Lshortfile)
	debuglog, err1 := os.OpenFile("debug.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err1 != nil {
		log.Println("打开日志文件错误")
	}
	errorlog, err2 := os.OpenFile("error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err2 != nil {
		log.Println(err2)
	}

	Info = log.New(os.Stdout,"Info:",log.Lmicroseconds)
	Debug = log.New(io.MultiWriter(debuglog, os.Stdout), "Debug:", log.LstdFlags|log.Lshortfile)
	Error = log.New(io.MultiWriter(errorlog, os.Stdout), "Error:", log.LstdFlags|log.Lshortfile)

}
