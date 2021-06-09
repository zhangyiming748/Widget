package unzip

import (
	"allTools/util/log"
	"os/exec"
	"strings"
)

/*
读取密码并尝试破解
*/
//7za x -ppassword filename.7z
//src 文件位置
//dst解压位置
//file解压密码(单条文本)
func Passwd(src, dst, password string) {
	p := strings.Join([]string{"-p", password}, "")
	o := strings.Join([]string{"-o", dst}, "")
	cmd := exec.Command("7za", "x", p, src, o)
	log.Info.Printf("生成的命令是:%s", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Info.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Info.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.CMD.Printf("%v", string(tmp))
		t := string(tmp)
		if s := t; strings.Contains(s, "replace") {
			log.Info.Println("文件已存在,跳过")
		}
		if s := t; strings.Contains(s, "Wrong password") {
			log.Info.Printf("当前密码%s错误,跳过", password)
		}
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.CMD.Println("运行中有错误产生", err)
	} else {
		log.Info.Printf("文件的密码有可能是%s\n", password)
	}

}
