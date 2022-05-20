package file

import (
	"Widget/util/log"
	"os/exec"
)

func RunShell(absPath string) {
	log.Info.Printf("即将运行的shell:%s\n", absPath)
	sh := exec.Command("bash", "-c", absPath)
	output, err := sh.Output()
	if err != nil {
		log.Warn.Panicf("运行shell发生错误:%v\n", err)
	} else {
		log.Info.Printf("%s\n", output)
	}
	stdout, _ := sh.StdoutPipe()
	sh.Stderr = sh.Stdout
	sh.Start()
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		log.Info.Printf("文件输出:%s\n", tmp)
		if err != nil {
			break
		}
	}
	sh.Wait()
}
