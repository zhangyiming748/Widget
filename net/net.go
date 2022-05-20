package net

import (
	"Widget/util/log"
	"os/exec"
)

//hey -n 2000 -c 50 -m GET http://127.0.0.1:8889/HappyCount
func Hey(URL, Requests, Concurrent string) {
	cmd := exec.Command("hey", "-n", Requests, "-c", Concurrent, "-m", "GET", URL)
	log.Info.Printf("生成的命令是:%s", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.CMD.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.CMD.Printf("cmd.Run产生的错误:%v", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Debug.Println(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.CMD.Println("命令执行中有错误产生", err)
	}
}
