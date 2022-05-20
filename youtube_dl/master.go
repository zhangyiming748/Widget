package youtube_dl

import (
	"Widget/util/file"
	"Widget/util/log"
	"errors"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Master(fp, addr, port, target, MaxGoroutine, isproxy string) {
	if fp == "" {
		panic(errors.New("没有找到待下载文件列表"))
	}
	if addr == "" {
		panic(errors.New("没有有效的ip"))
	}
	if port == "" {
		panic(errors.New("没有有效的端口"))
	}
	if target == "" {
		panic(errors.New("没有有效目标文件夹"))
	}
	if MaxGoroutine == "no value" || MaxGoroutine == "" {
		MaxGoroutine = "1"
		log.Info.Println("没有定义最大连接数,默认为单线程")
	}
	proxy := strings.Join([]string{addr, port}, ":")
	log.Info.Printf("youtube-dl代理设置为\n", proxy)
	var wg sync.WaitGroup
	max, _ := strconv.Atoi(MaxGoroutine)
	ch := make(chan struct{}, max)
	//ToDO
	//links := []string{} //readline.Readlink(fp)//读取全部文本文件到切片
	links := file.ReadLink(fp)
	list := make(map[string]bool)
	for i, v := range links {
		if list[v] == true {
			log.Info.Printf("跳过重复文件No.%d", i+1)
			continue
		}
		if v == "javascript:void(0)" {
			log.Info.Printf("跳过不存在的文件No.%d", i+1)
			continue
		}
		ch <- struct{}{}
		wg.Add(1)
		log.Info.Printf("开始尝试下载NO.%d\n", i+1)
		go RunCmd(v, &wg, proxy, target, i, ch, isproxy)
		list[v] = true
	}
	wg.Wait()
}
func RunCmd(url string, wg *sync.WaitGroup, proxy, dir string, i int, ch chan struct{}, isproxy string) {
	path := strings.Join([]string{dir, "%(title)s.%(ext)s"}, "/")
	s := strings.Join([]string{"%(title)s", "m3u8"}, ".")
	log.Info.Printf("得到的文件名: %s", s)
	var cmd *exec.Cmd
	if isproxy == "0" { //不使用代理
		cmd = exec.Command("youtube-dl", "-o", path, "-f", "best", url)
		log.Info.Println("不使用代理")
	}
	if isproxy == "1" { //使用代理
		cmd = exec.Command("youtube-dl", "--proxy", proxy, "-o", path, "-f", "best", url)
		log.Info.Println("使用代理")
	}

	log.Info.Printf("生成的命令是: %s", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Info.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Info.Printf("cmd.Run产生的错误:%v", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if strings.HasSuffix(string(tmp), "has already been downloaded") {
			log.Info.Println("当前文件已存在")
		}
		log.Info.Printf("第%d个文件输出:%s", i+1, string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Info.Printf("%v对应文件:%v\n", err, url)
	}
	log.Info.Printf("下载文件%v完成\n", url)
	wait := time.Duration(rand.Intn(3))
	time.Sleep(wait * time.Second)
	<-ch
	wg.Done()
}
