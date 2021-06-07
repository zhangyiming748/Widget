package main

import (
	"allTools/convert"
	"allTools/rotateVideos"
	conf "allTools/util/conf"
	util "allTools/util/file"
	"os"
	"path/filepath"
	"sync"
	"time"

	// f "allTools/util/file"
	"allTools/util/log"
)

var (
	files   []string
	pattern = conf.GetVal("location", "pattern")
	src     = conf.GetVal("location", "src")
	dst     = conf.GetVal("location", "dst")
)

func init() {
	dst := conf.GetVal("location", "dst")
	if exists(dst) || isDir(dst) {
		log.Info.Println("目标文件夹设置正确")
	} else {
		//os.Remove(dst)
		log.Info.Println("新建目标文件夹")
		os.Mkdir(dst, 0644)
	}
}
func initial(wg *sync.WaitGroup)  {
	files, err := filepath.Glob("./*.log")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
	time.Sleep(5*time.Second)
	wg.Done()
}

func main() {
	files = util.GetFiles(src, pattern)
	fn := conf.GetVal("main", "function")
	switch fn {
	case "ToMp4":
		log.Info.Println("ToMp4")
		for _, file := range files {
			convert.ToMp4(src, file)
		}
	case "rotate":
		log.Info.Println("rotate")
		if direct := conf.GetVal("rotate", "direction"); direct == "ToRight" {
			for _, file := range files {
				rotateVideos.ToRight(src, dst, file)
			}
		} else {
			for _, file := range files {
				rotateVideos.ToLeft(src, dst, file)
			}
		}
	case "ToWebp":
		log.Info.Println("ToWebp")
		for _, file := range files {
			convert.ToWebp(src, file)
		}
	case "WebpTo":
		log.Info.Println("WebpTo")
		for _, file := range files {
			convert.WebpTo(src, file)
		}
	}
}

// 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
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
