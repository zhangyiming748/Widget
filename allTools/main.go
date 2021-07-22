package main

import (
	"allTools/convert"
	"allTools/getFileType"
	"allTools/rotateVideos"
	"allTools/unzip"
	conf "allTools/util/conf"
	util "allTools/util/file"
	"allTools/util/log"
	"allTools/weather"
	"os"
	"strings"
)

var (
	files   []string
	pattern = conf.GetVal("location", "pattern")
	src     = conf.GetVal("location", "src")
	dst     = conf.GetVal("location", "dst")
	delDone  = conf.GetVal("main", "delAfterDone")
)

func init() {
	dst := conf.GetVal("location", "dst")
	if exists(dst) || isDir(dst) {
		log.Info.Println("目标文件夹设置正确")
	} else {
		//os.Remove(dst)
		log.Info.Println("新建目标文件夹")
		if err := os.Mkdir(dst, 0644); err != nil {
			/*

			 */
		}
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Info.Printf("程序运行过程中有错误产生:%v", err)
		}
	}()
	files = util.GetFiles(src, pattern)
	fn := conf.GetVal("main", "function")
	switch fn {
	case "ToMp4":
		log.Info.Println("ToMp4")
		for _, file := range files {
			log.Debug.Printf("准备好进行转换的文件:%v", file)
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
		log.Info.Println("webp最大宽高不得超过16383像素")
		for _, file := range files {
			convert.ToWebp(src, file)
		}
	case "WebpTo":
		log.Info.Println("WebpTo")
		for _, file := range files {
			convert.WebpTo(src, file)
		}
	case "Unzip":
		log.Info.Println("Unzip")
		keyfile := conf.GetVal("location", "passwd")
		passwords := util.ReadLine(keyfile)
		for _, passwd := range passwords {
			unzip.UnZip(src, dst, passwd)
		}
	case "Weather":
		log.Info.Println("查询天气")
		weather.Weather()
	case "Detect":
		log.Info.Println("探测文件类型")
		files := util.GetAllFiles(src)
		for _, file := range files {
			full := strings.Join([]string{src, file}, "/")
			log.Debug.Println(full)
			log.Info.Printf("(%s)的文件类型是:(%s)", file, getFileType.Detect(full))
		}
	case "ExtractAudio":
		log.Info.Println("提取视频中的音频")
		files := util.GetAllFiles(src)
		for i, file := range files {
			convert.ExtractAudio(src, file)
			log.Info.Printf("处理第(%d / %d)个文件\n",i+1,len(files))
			if delDone == "true" {
				s := strings.Join([]string{src, file}, "/")
				if err := os.Remove(s); err == nil {
					log.Info.Printf("删除转换完成的文件%v成功\n", s)
				} else {
					log.Info.Printf("删除转换完成的文件%v发生错误\n", s)
				}
			}
		}
	default:
		//log.Info.Println("")
		panic("没有指定程序功能")
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
