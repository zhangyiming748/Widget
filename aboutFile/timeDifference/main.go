package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	var (
		stime, etime string
	)
	fmt.Println("输入开始时间")
	fmt.Scanf("%s", &stime)
	fmt.Println("输入结束时间")
	fmt.Scanf("%s", &etime)
	sec := subTime(stime, etime)
	keeptime := makeHHMMSS(int(sec))
	fmt.Printf("起始时间%v持续时间%v", stime, keeptime)
}

/*
将输出的秒转换为hh:mm:ss
*/
func makeHHMMSS(sec int) string {

	hh := sec % 3600
	sec = sec - 3600*hh
	mm := sec % 60
	sec = sec - 60*mm
	ss := sec
	stime := strings.Join([]string{string(hh), string(mm), string(ss)}, ":")
	//fmt.Printf("持续时间%s\n", stime)
	return stime
}

/*
计算时间差，输出秒
*/
func subTime(start, end string) int64 {
	timeLayout :="2006-01-02 03:04:05 PM"
	loc, _ := time.LoadLocation("Local")
	s_time, err := time.ParseInLocation(timeLayout, start, loc)
	if err != nil {
		log.Println("StringToTimeStamp出现异常：", err)
	}
	e_time, err := time.ParseInLocation(timeLayout, end, loc)
	if err != nil {
		log.Println("StringToTimeStamp出现异常：", err)
	}

	start_time := s_time.Unix()
	end_time := e_time.Unix()
	t := end_time - start_time
	return t
}
