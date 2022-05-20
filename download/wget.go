package download

import (
	"Widget/util/log"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

func WGet(url, dst string) {
	cmd := exec.Command("wget", url, "-P", dst)
	log.Info.Println("不使用代理")
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
		if strings.Contains(string(tmp), "s") && strings.Contains(string(tmp), "%") {
			log.Info.Println(string(tmp))
		}
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
}
