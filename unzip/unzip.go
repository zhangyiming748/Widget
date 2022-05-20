package unzip

import (
	"Widget/util/log"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

/*
读取密码并尝试破解
*/
//7za x -ppassword filename.7z
//src 文件位置
//dst解压位置
//file解压密码(单条文本)
func UnZip(src, dst, password string, done chan string) {
	p := strings.Join([]string{"-p", password}, "")
	o := strings.Join([]string{"-o", dst}, "")
	cmd := exec.Command("7za", "x", p, src, o)
	log.Debug.Printf("生成的命令是:%s", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Debug.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Debug.Printf("需要使用7za命令,通过 sudo apt-get install p7zip-full 安装")
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.CMD.Printf("%v", string(tmp))
		t := string(tmp)
		if s := t; strings.Contains(s, "replace") {
			log.Debug.Println("文件已存在,跳过")
		}
		if s := t; strings.Contains(s, "Wrong password") {
			log.Debug.Printf("当前密码%s错误,跳过", password)
		}
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		//log.CMD.Println("此次运行密码错误", err)
		delFail(dst)
	} else {
		//log.Info.Printf("文件的密码有可能是%s\n", password)
		done <- password
	}

}

//删除已经解压但用了错误密码的空文件
func delFail(dst string) {
	dir, _ := ioutil.ReadDir(dst)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{dst, d.Name()}...))
	}

}
