package util

import (
	"fmt"
	"os/exec"
)

func Alarm() {
	//echo -e "\a"
	cmd := exec.Command("echo", "-e", "\"\\a\"")
	// 命令的错误输出和标准输出都连接到同一个管道
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd.Run产生的错误:%v", err)
	}
}
