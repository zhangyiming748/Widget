package downloadcmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"youtube-dl/mylog"
)

//youtube-dl -o "~/Desktop/%(title)s.%(ext)s" 'youtube file url'
//func RunCommand(url string, wg *sync.WaitGroup, i int) {
//	cmd := exec.Command("youtube-dl", "--proxy", "127.0.0.1:8889", "-o", "/Users/zen/Downloads/trans/%(title)s.%(ext)s", "-f", "best", url)
//	var stdout, stderr bytes.Buffer
//	cmd.Stdout = &stdout
//	cmd.Stderr = &stderr
//	err := cmd.Run()
//	if err != nil {
//		log.Fatalf("cmd.Run() failed with %s\n", err)
//	}
//	//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
//	//fmt.Printf("Str:\n%s\nerr:\n%s\n", outStr, errStr)
//	fmt.Printf("下载NO.%d完成\n", i)
//	wg.Done()
//}
func RunCmd(url string, wg *sync.WaitGroup, proxy, dir string) {
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
		log.Printf("输出:%s", string(tmp))
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
	wg.Done()
}
func split(s string) string {
	strs := strings.Split(s, "/")
	suffix := strs[len(strs)-1]
	return suffix
}
