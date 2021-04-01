package util

import (
	"io/ioutil"
	"strings"
)
/**
读取配置中的指定文件夹并返回绝对路径
 */
func GetFiles()(flist []string){
	dir:=GetVal("src","fp")
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		//fmt.Println(f.Name())
		if strings.HasSuffix(f.Name(),GetVal("format","pattern")){
			if !strings.HasPrefix(f.Name(),"."){
				absPath:=strings.Join([]string{dir,f.Name()},"/")
				flist=append(flist,absPath)
			}
		}
	}
	return
}
