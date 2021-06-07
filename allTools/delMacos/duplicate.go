package delMacos

import (
	"allTools/util/log"
	"io/ioutil"
	"os"
	"strings"
)
//删除多余的文件挂载点
func delFiles(dir string) []string {
	files, _ := ioutil.ReadDir(dir)
	aim := []string{}
	for _, f := range files {
		if f.IsDir() { //如果是目录就跳过
			continue
		}
		log.Info.Println(f.Size())
		if name := f.Name(); strings.HasPrefix(name, ".")  {
			log.Info.Printf("即将删除的文件%s\n", name)
			if err := os.Remove(name); err != nil {
				log.Info.Printf("删除%s失败\n", name)
			}
		}
		if name := f.Name(); name == ".DS_Store" {
			log.Info.Printf("即将删除的文件%s\n", name)
			if err := os.Remove(name); err != nil {
				log.Info.Printf("删除%s失败\n", name)
			}
		}

	}
	return aim
}
