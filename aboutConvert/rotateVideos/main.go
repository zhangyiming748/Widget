package main

import (
	"fmt"
	l "log"
	"os"
	"rotateVideos/shellCmd"
	"rotateVideos/util"
	"time"
)
/*
目前只能正确处理不带空格的文件名
 */
func init()  {
	dst := util.GetVal("target", "dst")
	if exists(dst)||isDir(dst){
		l.Println("目标文件夹设置正确")
	}else {
		l.Println("新建目标文件夹")
		os.Mkdir(dst,0777)
	}

}
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
	l.Printf("处理%d个文件,共用时%v\n",len(files),take)
}
// 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}


