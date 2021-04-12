package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		stime, etime string
	)
	fmt.Println("输入开始时间,格式HHMMSS")
	fmt.Scanf("%s", &stime)
	sh:=stime[0:2]
	sm:=stime[2:4]
	ss:=stime[4:]
	s_sec:=toSec(sh,sm,ss)
	fmt.Printf("开始时间:%v",strings.Join([]string{sh,sm,ss},":"))
	fmt.Println("输入结束时间,格式HHMMSS")
	fmt.Scanf("%s", &etime)
	eh:=etime[0:2]
	em:=etime[2:4]
	es:=etime[4:]
	e_sec:=toSec(eh,em,es)
	t:=e_sec-s_sec
	fmt.Printf("持续时间:%v",strings.Join([]string{eh,em,es},":"))


}
func toSec(hh,mm,ss string)int  {
	var sec int
	h,_:=strconv.Atoi(hh)
	m,_:=strconv.Atoi(mm)
	s,_:=strconv.Atoi(ss)
	sec=h*3600 +m*60+s
	return sec
}
func int2time(i int) string {
	h:=i/3600
	if h<10{
		h="0"+string(h)
	}
	t:= strings.Join([]string{string(),string(),string()},":")
	return ""
}