package convert

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
	. "youtube-dl/mylog"
)

func Convert(fullpath string) {
	toMp4(fullpath)
}
func toMp4(fullpath string) {
	out := strings.Join([]string{strings.Split(fullpath, ".")[0], "mp4"}, ".")
	var threads string
	var ffmpegCmd = []string{}
	if goos := runtime.GOOS; goos == "darwin" {
		threads = "1"
		ffmpegCmd = append(ffmpegCmd, "ffmpeg")
		ffmpegCmd = append(ffmpegCmd, "-threads")
		ffmpegCmd = append(ffmpegCmd, "1")
		ffmpegCmd = append(ffmpegCmd, "-i")
		ffmpegCmd = append(ffmpegCmd, "\""+fullpath+"\"")
		ffmpegCmd = append(ffmpegCmd, "-threads")
		ffmpegCmd = append(ffmpegCmd, "1")
		ffmpegCmd = append(ffmpegCmd, "\""+out+"\"")
		ffmpegCmd = append(ffmpegCmd, "|")
		ffmpegCmd = append(ffmpegCmd, "tee")
		ffmpegCmd = append(ffmpegCmd, "ffmpeg.log")
	} else {
		threads=string(runtime.NumCPU())
		ffmpegCmd = append(ffmpegCmd, "ffmpeg")
		ffmpegCmd= append(ffmpegCmd, "-threads")
		ffmpegCmd= append(ffmpegCmd, threads)
		ffmpegCmd = append(ffmpegCmd, "-i")
		ffmpegCmd = append(ffmpegCmd, "\""+fullpath+"\"")
		ffmpegCmd= append(ffmpegCmd, "-threads")
		ffmpegCmd= append(ffmpegCmd, threads)
		ffmpegCmd = append(ffmpegCmd, "\""+out+"\"")
		ffmpegCmd = append(ffmpegCmd, "|")
		ffmpegCmd = append(ffmpegCmd, "tee")
		ffmpegCmd = append(ffmpegCmd, "ffmpeg.log")
	}
	var cmdInLine string
	for _, v := range ffmpegCmd {
		cmdInLine = strings.Join([]string{cmdInLine, v}, " ")
	}
	Debug.Printf("前半部分命令:%s", cmdInLine)
	Debug.Printf("当前使用的线程数是%v", threads)
	cmd := exec.Command("bash", "-c", cmdInLine)
	Info.Printf("生成的命令是:%s", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		Error.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		Error.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		Debug.Println(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Println("命令执行中有错误产生", err)
	}
}
