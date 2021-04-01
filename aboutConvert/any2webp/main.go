package main

import (
	"any2webp/command"
	"any2webp/util"
	"fmt"
	"sync"
)

/**
WebP支援的像素最大數量是16383x16383
*/
func main() {
	files := util.GetFiles()
	ch := make(chan struct{}, 3)
	fmt.Println(files)
	var wg sync.WaitGroup
	for i, f := range files {
		wg.Add(1)
		go command.RunCmd(f, &wg, i, ch)
	}
	wg.Wait()
}
