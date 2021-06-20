package main

import (
	"fmt"
	"v264add2webm/shell"
	"v264add2webm/util"
)

func main() {
	dir := util.GetVal("target", "src")
	ret := util.GetFiles(dir)
	fmt.Printf("符合条件的文件是:%s\n", ret)
	util.MakeList(ret)
	shell.Shell()
}
