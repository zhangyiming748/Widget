package main

import (
	"fmt"
	"rotateVideos/shellCmd"
	"rotateVideos/util"
)
/*
目前只能正确处理不带空格的文件名
 */
func main() {
	var (
		param string
		files []string
		src = util.GetVal("target", "src")
		dst = util.GetVal("target", "dst")
	)
	files = util.GetFiles(src, util.GetVal("target", "pattern"))
	fmt.Println(files)
	param = util.GetVal("direction", "before")
	if param == "left"{
		for _,file:=range files{
			shellCmd.LeftTo(src,dst,file)
		}
	}else if param=="right"{
		for _,file:=range files{
			shellCmd.RightTo(src,dst,file)
		}
	}else {
		panic("旋转参数未定义")
	}




}
