package main


import (
	"fmt"
	"sync"
	"youtube-dl/cmd"
	"youtube-dl/readline"
)

func main() {
	fp := "links.txt"
	var wg sync.WaitGroup
	links := readline.Readlink(fp)
	for i, v := range links {
		wg.Add(1)
		fmt.Printf("开始尝试下载NO.%d\n", i)
		go cmd.Ytd(v, &wg, i)
	}
	wg.Wait()
}
