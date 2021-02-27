package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

//youtube-dl -o "~/Desktop/%(title)s.%(ext)s" 'youtube file url'
func RunCommand(url string, wg *sync.WaitGroup, i int) {
	cmd := exec.Command("youtube-dl", "--proxy", "127.0.0.1:8889", "-o", "/Users/zen/Downloads/trans/%(title)s.%(ext)s", "-f", "best", url)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Printf("Str:\n%s\nerr:\n%s\n", outStr, errStr)
	fmt.Printf("下载NO.%d完成\n", i)
	wg.Done()
}
func RunCmd(url string, wg *sync.WaitGroup, i int) {
	cmd := exec.Command("youtube-dl", "--proxy", "127.0.0.1:8889", "-o", "/Users/zen/Downloads/trans/%(title)s.%(ext)s", "-f", "best", url)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		fmt.Printf("cmd.StdoutPipe产生的错误:%v",err)
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("cmd.Run产生的错误:%v",err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Printf("NO.%d的Goroutine输出:%s\n",i,string(tmp))
		if err != nil {
			break
		}
	}
	fmt.Printf("下载NO.%d完成\n", i)
	if err = cmd.Wait(); err != nil {
		fmt.Println(err)
	}
	wg.Done()
}
