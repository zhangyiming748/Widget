package main

import (
	"any2mp4/shellCmd"
	"any2mp4/util"
	"fmt"
)

/*
目前只能正确处理不带空格的文件名
*/
func main() {
	var (
		files []string
		src   = util.GetVal("target", "src")
	)
	files = util.GetFiles(src)
	fmt.Println(files)
	for _, file := range files {
		shellCmd.ToMp4(src, file)
	}

}
