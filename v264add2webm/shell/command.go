package shell

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"v264add2webm/util"
)

//ffmpeg -f concat -safe 0 -i work.txt -c copy /Users/zen/Movies/out.avi
///Users/zen/Github/Widget/v264add2webm/list.txt
func Shell() error {
	txt := util.GetVal("target", "list")
	dst := util.GetVal("target", "dst")
	output := strings.Join([]string{dst, "mp4"}, ".")
	//output := strings.Join([]string{strings.Split(txt, ".")[0], "webm"}, ".")
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", txt, "-c", "copy", output)
	fmt.Printf("生成的命令是:%s", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		fmt.Printf("cmd.StdoutPipe产生的错误:%v", err)
		return err
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("cmd.Run产生的错误:%v", err)
		return err
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Println(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Println("命令执行中有错误产生", err)
		return err
	}
	return nil
}
