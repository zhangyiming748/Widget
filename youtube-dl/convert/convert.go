package convert

import (
	"io/ioutil"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	. "youtube-dl/mylog"
	"youtube-dl/util"
)

func getFiles(dir string) []string {
	files, _ := ioutil.ReadDir(dir)
	aim := []string{}
	for _, f := range files {
		//fmt.Println(f.Name())
		if l := strings.Split(f.Name(), ".")[0]; len(l) != 0 {
			Debug.Printf("目录中有效的文件是:%v\n", f.Name())
			if strings.HasSuffix(f.Name(), "m3u8") || strings.HasSuffix(f.Name(), "webm") {
				FFmpeg.Printf("有效的目标文件是:%v\n", f.Name())
				aim = append(aim, f.Name())
			}
		}
	}
	return aim
}

func Convert() {
	dir := util.GetVal("target", "fp")
	Debug.Printf("读取到的目录是:%s", dir)
	files := getFiles(dir)
	for _, file := range files {
		toMp4(dir, file)
	}
}

func toMp4(src, file string) {
	var threads int = runtime.NumCPU()
	if system := runtime.GOOS; system == "darwin" {
		threads = 1
	} else {
		threads = 4
	}
	in := strings.Join([]string{src, file}, "/")
	fname := strings.Split(file, ".")[0]
	fname = strings.Join([]string{fname, "mp4"}, ".")
	out := strings.Join([]string{src, fname}, "/")
	t := strconv.Itoa(threads)
	cmd := exec.Command("ffmpeg", "-threads", t, "-i", in, "-threads", t, out)
	FFmpeg.Printf("生成的命令是:%s", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		FFmpeg.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		FFmpeg.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		Info.Println(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		FFmpeg.Println("命令执行中有错误产生", err)
	}
}
