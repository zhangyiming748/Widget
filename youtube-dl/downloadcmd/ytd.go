package downloadcmd

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"sync"
	"time"
	"youtube-dl/mylog"
)

func RunCmd(url string, wg *sync.WaitGroup, proxy, dir string, i int, ch chan struct{}) {
	path := strings.Join([]string{dir, "%(title)s.%(ext)s"}, "/")
	cmd := exec.Command("youtube-dl", "--proxy", proxy, "-o", path, "-f", "best", url)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		fmt.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("cmd.Run产生的错误:%v", err)
	}
	fn := split(url)
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Printf("第%d个文件输出:%s", i+1, string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		ret := fmt.Sprintf("命令运行期间产生的错误:%v\t对应文件:%v\n", err, url)
		mylog.Logof(ret)
		//log.Printf("重试下载%v\n", fn)
		//wg.Add(1)
		//time.Sleep(3 * time.Second)
		//go RunCmd(url, wg, proxy, dir)
	}
	ret := fmt.Sprintf("下载文件%v完成\n", fn)
	mylog.Logof(ret)
	wait:=time.Duration(rand.Intn(3))
	time.Sleep(wait*time.Second)
	<-ch
	wg.Done()
}
func split(s string) string {
	strs := strings.Split(s, "/")
	suffix := strs[len(strs)-1]
	return suffix
}
