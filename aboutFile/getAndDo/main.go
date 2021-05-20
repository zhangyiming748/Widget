package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	pattern = "apk"
)

func main() {
	files := GetFileList()
	for _, v := range files {
		command(v)
	}
}
func GetFileList() []string {
	pwd, _ := os.Getwd()
	//pattern := "go"
	flist := make([]string, 0)
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		//fmt.Println(fileInfoList[i].Name()) //打印当前文件或目录下的文件或目录名
		fname := fileInfoList[i].Name()
		//fmt.Printf("fname = %s\n", fname)
		if IsFile(fname, pattern) {
			flist = append(flist, fname)
		}
	}
	return flist
}
func IsFile(s, pattern string) bool {
	fs := strings.Split(s, ".")
	if len(fs) == 2 {
		if fs[1] == pattern {
			return true
		}
	}

	return false
}

func command(f string) {
	cmd := exec.Command("adb", "install", f)
	fmt.Printf("生成的命令是%v\n", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		fmt.Printf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("cmd.Run产生的错误:%v\n", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Printf("程序 %s 输出 %s\n", f, string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		fmt.Printf("命令运行期间产生的错误:%v\n", err)

	}
	fmt.Printf("安装完成\n")

}
