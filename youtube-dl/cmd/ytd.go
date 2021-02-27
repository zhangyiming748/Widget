package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync"
)
//youtube-dl -o "~/Desktop/%(title)s.%(ext)s" 'youtube file url'
func Ytd(url string,wg *sync.WaitGroup,i int) {
	cmd := exec.Command("youtube-dl", "--proxy","127.0.0.1:8889","-o","/Users/zen/Downloads/trans/%(title)s.%(ext)s","-f","best",url)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Printf("Str:\n%s\nerr:\n%s\n", outStr, errStr)
	fmt.Printf("下载NO.%d完成\n",i)
	wg.Done()
}
