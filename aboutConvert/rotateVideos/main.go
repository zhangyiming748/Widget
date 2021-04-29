package main

import (
	"fmt"
	"log"
	"rotateVideos/shellCmd"
	"rotateVideos/util"
	"time"
)
/*
目前只能正确处理不带空格的文件名
 */
func main() {
	start:=time.Now()
	var (
		param string
		files []string
		src = util.GetVal("target", "src")
		dst = util.GetVal("target", "dst")
	)
	files = util.GetFiles(src, util.GetVal("target", "pattern"))
	fmt.Printf("符合条件的文件: %v\n",files)
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
	end:=time.Now()
	take:=end.Sub(start)
	log.Printf("处理%d个文件,共用时%v\n",len(files),take)
}
