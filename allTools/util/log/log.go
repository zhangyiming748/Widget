package log

import (
	conf "allTools/util/conf"
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	CMD   *log.Logger
	Debug *log.Logger
)

func init() {
	prefix := conf.GetVal("main", "function")
	log.SetPrefix(prefix)
	log.SetFlags(log.Ltime | log.Lshortfile)
	cmdlog, err1 := os.OpenFile("cmd.log", os.O_WRONLY|os.O_CREATE, 0666)

	if err1 != nil {
		log.Println("打开日志文件错误")
	}
	infolog, err2 := os.OpenFile("info.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		log.Println("打开日志文件错误")
	}

	Info = log.New(io.MultiWriter(infolog, os.Stdout), prefix, log.LstdFlags|log.Lshortfile)
	CMD = log.New(io.MultiWriter(cmdlog, os.Stdout), prefix, log.LstdFlags|log.Lshortfile)
	//Error = log.New(io.MultiWriter(errorlog, os.Stdout), "Error:", log.LstdFlags|log.Lshortfile)
	Debug = log.New(os.Stdout, "Debug:", log.LstdFlags|log.Lshortfile)
}
