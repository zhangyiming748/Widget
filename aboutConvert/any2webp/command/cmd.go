package command

import (
	"fmt"
	"log"

	//"fmt"
	//"log"
	"os/exec"
	"strings"
	"sync"
)

func RunCmd(fp string, wg *sync.WaitGroup, i int, ch chan struct{}) {
	path := strings.Split(fp, ".")
	target := strings.Join([]string{path[0], "webp"}, ".")
	cmd := exec.Command("cwebp", "-lossless", "-z", "9", "-mt", "-sharp_yuv", "-v", fp, "-o", target)
	//命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		fmt.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("cmd.Run产生的错误:%v", err)
	}
	//fn := split(url)
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Printf("第%d个文件输出:%s", i, string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Printf("命令运行期间产生的错误")
	}
	<-ch
	wg.Done()
}
