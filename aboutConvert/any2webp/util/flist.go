package util

import (
	"io/ioutil"
	"strings"
)

/**
读取配置中的指定文件夹并返回绝对路径
*/
func GetFiles() (flist []string) {
	dir := GetVal("src", "fp")
	files, _ := ioutil.ReadDir(dir)
	if GetVal("runmode","mode")=="toWebp"{
		for _, f := range files {
			//fmt.Println(f.Name())
			if strings.HasSuffix(f.Name(), ".jpg") {
				if !strings.HasPrefix(f.Name(), ".") {
					absPath := strings.Join([]string{dir, f.Name()}, "/")
					flist = append(flist, absPath)
				}
			}
		}
	}
	if GetVal("runmode","mode")=="webpTo"{
		for _, f := range files {
			//fmt.Println(f.Name())
			if strings.HasSuffix(f.Name(), ".webp") {
				if !strings.HasPrefix(f.Name(), ".") {
					absPath := strings.Join([]string{dir, f.Name()}, "/")
					flist = append(flist, absPath)
				}
			}
		}
	}

	return
}
