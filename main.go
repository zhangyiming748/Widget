package main

import (
	convert "Widget/convert"
	"Widget/download"
	"Widget/net"
	"Widget/rotateVideo"
	t "Widget/time"
	"Widget/unzip"
	"Widget/util/conf"
	"Widget/util/file"
	"Widget/util/file/hash"
	"Widget/util/log"
	"Widget/weather"
	ytd "Widget/youtube_dl"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var (
	files        []string
	pattern      = conf.GetVal("location", "pattern")
	src          = conf.GetVal("location", "src")
	dst          = conf.GetVal("location", "dst")
	delDone      = conf.GetVal("main", "delAfterDone")
	MaxGoroutine = conf.GetVal("youtube_dl", "goroutine")
	fp           = conf.GetVal("youtube_dl", "fp")
	addr         = conf.GetVal("youtube_dl", "addr")
	port         = conf.GetVal("youtube_dl", "port")
	target       = conf.GetVal("youtube_dl", "target")
	proxy        = conf.GetVal("youtube_dl", "isproxy")
	URL          = conf.GetVal("hey", "url")
	Requests     = conf.GetVal("hey", "Requests")
	Concurrent   = conf.GetVal("hey", "Concurrent")
	OS           = runtime.GOOS
	ARCH         = runtime.GOARCH
	videoPrefix  = conf.GetVal("title", "start")
	multiLinks   = conf.GetVal("download", "multi")
)

func init() {
	initialization := exec.Command("bash", "-c", "/Users/zen/Github/Tools/init.sh")
	if b, err := initialization.Output(); err != nil {
		log.Warn.Panicln(err)
	} else {
		log.Info.Printf("\n初始化产生的结果:%s\n", b)
	}
	log.Info.Printf("当前系统是%s\t架构是%s\n", OS, ARCH)
}
func main() {
	cmd := exec.Command("/bin/bash", "-c", "./init.sh|tee ./init.log")
	if err := cmd.Start(); err != nil {
		log.Info.Printf("\"init.sh\"产生的错误:%v", err)
	}
	if err := cmd.Wait(); err != nil {
		log.CMD.Println("\"init.sh\"执行中有错误产生", err)
	}
	start := time.Now()
	log.Info.Println("程序开始时间:", time.Now().Format("2006-01-02 15:04:05"))
	defer func() {
		end := time.Now()
		log.Info.Println("程序结束时间:", time.Now().Format("2006-01-02 15:04:05"))
		sub := end.Sub(start)
		log.Info.Println("程序用时:", sub)
	}()
	defer func() {
		if err := recover(); err != nil {
			log.Info.Printf("程序运行过程中有错误产生:%v", err)
		}
	}()

	files = file.GetFiles(src, pattern)
	fn := conf.GetVal("main", "function")
	switch fn {
	case "ToMp4":
		log.Info.Println("ToMp4")
		for _, f := range files {
			log.Info.Printf("准备好进行转换的文件:%v", f)
			convert.ToMp4(src, f)
			if delDone == "true" {
				s := strings.Join([]string{src, f}, string(os.PathSeparator))
				if err := os.Remove(s); err == nil {
					log.Info.Printf("删除转换完成的文件%v成功\n", s)
				} else {
					log.Info.Printf("删除转换完成的文件%v发生错误\n", s)
				}
			}
		}
	case "ToWebm":
		log.Info.Println("ToWebm")
		for _, f := range files {
			log.Info.Printf("准备好进行转换的文件:%v", f)
			convert.ToWebm(src, f)
			if delDone == "true" {
				s := strings.Join([]string{src, f}, string(os.PathSeparator))
				if err := os.Remove(s); err == nil {
					log.Info.Printf("删除转换完成的文件%v成功\n", s)
				} else {
					log.Info.Printf("删除转换完成的文件%v发生错误\n", s)
				}
			}
		}
	case "rotate":
		log.Info.Println("rotate")
		if direct := conf.GetVal("rotate", "direction"); direct == "ToRight" {
			for _, f := range files {
				log.Info.Printf("准备好进行转换的文件:%v", f)
				rotateVideo.ToRight(src, dst, f)
				if delDone == "true" {
					s := strings.Join([]string{src, f}, string(os.PathSeparator))
					if err := os.Remove(s); err == nil {
						log.Info.Printf("删除转换完成的文件%v成功\n", s)
					} else {
						log.Info.Printf("删除转换完成的文件%v发生错误\n", s)
					}
				}
			}
		} else {
			for _, f := range files {
				log.Info.Printf("准备好进行转换的文件:%v", f)
				rotateVideo.ToLeft(src, dst, f)
				if delDone == "true" {
					s := strings.Join([]string{src, f}, string(os.PathSeparator))
					if err := os.Remove(s); err == nil {
						log.Info.Printf("删除转换完成的文件%v成功\n", s)
					} else {
						log.Info.Printf("删除转换完成的文件%v发生错误\n", s)
					}
				}
			}
		}
	case "ToWebp":
		log.Info.Println("webp最大宽高不得超过16383像素")
		for _, f := range files {
			convert.ToWebp(src, f)
		}
	case "WebpTo":
		log.Info.Println("WebpTo")
		for _, f := range files {
			convert.WebpTo(src, f)
		}
	case "ToGif":
		log.Info.Println("ToGif")
		for _, f := range files {
			convert.ToGif(src, f)
		}
	case "Unzip":
		log.Info.Println("Unzip")
		keyfile := conf.GetVal("unzip", "passwd")
		crack := make(chan string, 1)
		passwords := file.ReadLine(keyfile)
		go func() {
			for _, passwd := range passwords {
				unzip.UnZip(src, dst, passwd, crack)
			}
		}()
		if v, ok := <-crack; ok {
			log.Info.Printf("密码有可能是%v", v)
			break
		}
	case "wget":
		log.Info.Println("download f")
		list := conf.GetVal("download", "wget")
		files := file.ReadLink(list)
		for _, f := range files {
			download.WGet(f, dst)
		}

	case "Weather":
		log.Info.Println("查询天气")
		weather.Weather()
	case "ExtractAudio":
		log.Info.Println("提取视频中的音频")
		files := file.GetAllFiles(src)
		for i, f := range files {
			convert.ExtractAudio(src, f)
			log.Info.Printf("处理第(%d / %d)个文件\n", i+1, len(files))
			if delDone == "true" {
				s := strings.Join([]string{src, f}, string(os.PathSeparator))
				if err := os.Remove(s); err == nil {
					log.Info.Printf("删除转换完成的文件%v成功\n", s)
				} else {
					log.Info.Printf("删除转换完成的文件%v发生错误\n", s)
				}
			}
		}
	case "GetTime":
		log.Info.Println("获取时间差(ffmpeg的-t)")
		t.TimeDiff()
	case "youtube_dl":
		log.Info.Println("youtube-dl下载")
		ytd.Master(fp, addr, port, target, MaxGoroutine, proxy)
	case "hey":
		log.Info.Println("压力测试网页")
		net.Hey(URL, Requests, Concurrent)
	case "WeatherPNG":
		log.Info.Println("生成天气预报图片")
	case "HASH":
		if isDir(src) {
			files = file.GetFiles(src, pattern)
			hash.SHA1(files...)
			hash.SHA256(files...)
			hash.MD5(files...)
		} else {
			if sameFile(src, dst) {
				log.Info.Printf("%v和%v是同一个文件\n", src, dst)
			} else {
				log.Info.Printf("%v和%v不是同一个文件\n", src, dst)
			}

		}
	case "duplicate":
		file.Duplicate(src, dst)
	case "decode":
		files = file.GetFiles(src, pattern)
		for _, f := range files {
			fullpath := strings.Join([]string{src, f}, string(os.PathSeparator))
			log.Info.Printf("文件路径为%v\n", f)
			convert.GB18030ToUtf8(fullpath)
		}
	case "cut":
		files = file.GetFiles(src, pattern)
		for _, f := range files {
			convert.Cut(src, f, videoPrefix)
		}
	case "multi":
		links := file.ReadLink(multiLinks)
		for _, link := range links {
			url := strings.Split(link, "|")[0]
			name := strings.Split(link, "|")[1]
			log.Info.Printf("获取到的下载链接:%s\n保存的文件名是%s\n", url, name)
			err := download.MultiDownload(url, name, runtime.NumCPU(), true)
			if err != nil {
				continue
			}
		}
	case "ToFlac":
		files := file.GetFiles(src, pattern)
		for _, f := range files {
			convert.ToFlac(src, f)
		}
	case "ToMp3":
		files := file.GetFiles(src, pattern)
		for _, f := range files {
			convert.ToMp3(src, f)
		}
	case "default":
		log.Info.Println("不运行功能")
	default:
		panic("没有指定程序功能或错误的拼写")
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
func sameFile(src, dst string) bool {
	s1f1, _ := hash.SHA1File(src)
	s256f1, _ := hash.SHA256File(src)
	md5f1, _ := hash.MD5File(src)
	s1f2, _ := hash.SHA1File(dst)
	s256f2, _ := hash.SHA256File(dst)
	md5f2, _ := hash.MD5File(dst)
	if s1f1 == s1f2 && s256f1 == s256f2 && md5f1 == md5f2 {
		return true
	}
	return false
}
