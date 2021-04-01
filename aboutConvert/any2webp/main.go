package main

import (
	"any2webp/command"
	"any2webp/util"
	"fmt"
	"log"
	"strconv"
	"sync"
)

/**
WebP支援的像素最大數量是16383x16383
*/
var (
	MaxGoroutine = util.GetVal("goroutine", "num")
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if val := util.GetVal("runmode", "mode"); val == "toWebp" {
		jpg2webp()
	} else if val == "webpTo" {
		webp2jpg()
	} else {
		panic("runmode设置错误")
	}
}

func jpg2webp() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	files := util.GetFiles()
	max, _ := strconv.Atoi(MaxGoroutine)
	ch := make(chan struct{}, max)
	fmt.Println(files)
	var wg sync.WaitGroup
	for i, f := range files {
		ch <- struct{}{}
		wg.Add(1)
		go command.ToWebp(f, &wg, i, ch)
	}
	wg.Wait()
}
func webp2jpg() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	files := util.GetFiles()
	max, _ := strconv.Atoi(MaxGoroutine)
	ch := make(chan struct{}, max)
	fmt.Println(files)
	var wg sync.WaitGroup
	for i, f := range files {
		ch <- struct{}{}
		wg.Add(1)
		go command.WebpTo(f, &wg, i, ch)
	}
	wg.Wait()
}
