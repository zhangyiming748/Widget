package downloadcmd

import (
	"math/rand"
	"os/exec"
	"strings"
	"sync"
	"time"
	"youtube-dl/convert"
	. "youtube-dl/mylog"
)

func RunCmd(url string, wg *sync.WaitGroup, proxy, dir string, i int, ch chan struct{}) {
	path := strings.Join([]string{dir, "%(title)s.%(ext)s"}, "/")
	cmd := exec.Command("youtube-dl", "--proxy", proxy, "-o", path, "-f", "best", url)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		Error.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		Error.Printf("cmd.Run产生的错误:%v", err)
	}
	//fn := split(url)
	//fn := url
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		Debug.Printf("第%d个文件输出:%s", i+1, string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		Error.Printf("%v对应文件:%v\n", err, url)
		//log.Printf("重试下载%v\n", fn)
		//wg.Add(1)
		//time.Sleep(3 * time.Second)
		//go RunCmd(url, wg, proxy, dir)
	}
	Info.Printf("下载文件%v完成\n", url)
	if ext:=strings.Split(path,".")[1];ext=="m3u8"{
		Debug.Printf("开始转换文件:%s",path)
		convert.ToMp4(path)
		Debug.Printf("转换完成文件:%s",path)
	}
	Debug.Println("当前文件处理结束")
	wait := time.Duration(rand.Intn(3))
	time.Sleep(wait * time.Second)
	<-ch
	wg.Done()
}
func getFormat(url string) {

}