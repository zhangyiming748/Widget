package util

import (
	"io/ioutil"
	"strings"
)

func GetFiles(dir string) []string {
	files, _ := ioutil.ReadDir(dir)
	aim := []string{}
	for _, f := range files {
		//fmt.Println(f.Name())
		if l := strings.Split(f.Name(), ".")[0]; len(l) != 0 {
			//fmt.Printf("有效的文件%v\n",f.Name())
			if strings.HasSuffix(f.Name(), "mkv") || strings.HasSuffix(f.Name(), "webm") {
				//fmt.Printf("有效的目标文件%v\n",f.Name())
				aim = append(aim, f.Name())
			}
		}
	}
	return aim
}
