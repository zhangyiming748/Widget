package main

import (
	"ArchivePasswd/rarArchive"
	"ArchivePasswd/readPasswd"
	"fmt"
	"os"
)

var (
	passwds = "./passwordDictionary.txt"
)

func main() {
	var file string
	if len(os.Args) == 2 {
		file = os.Args[1]
		//fmt.Println("文件名",file)
	} else {
		fmt.Println("没有指定文件,程序退出")
		os.Exit(0)
	}
	passwords := readPasswd.ReadPasswd2Slice(passwds)
	for _,ps:=range passwords{
		rarArchive.UnRar(file,ps)
	}
}
