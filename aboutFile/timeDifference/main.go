package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		stime, etime string
	)
	fmt.Printf("输入开始时间,格式HHMMSS\n")
	fmt.Scanf("%s", &stime)
	sh := stime[0:2]
	sm := stime[2:4]
	ss := stime[4:]
	s_sec := toSec(sh, sm, ss)
	hint,_:=strconv.Atoi(sh)
	mint,_:=strconv.Atoi(sm)
	sint,_:=strconv.Atoi(ss)
	//fmt.Printf("开始时间:%v\n", strings.Join([]string{sh, sm, ss}, ":"))
	fmt.Printf("输入结束时间,格式HHMMSS\n")
	fmt.Scanf("%s", &etime)
	eh := etime[0:2]
	em := etime[2:4]
	es := etime[4:]
	e_sec := toSec(eh, em, es)
	t := e_sec - s_sec
	kh,km,ks:=secTo(t)
	//h:=strconv.Itoa(sh)
	//m:=strconv.Itoa(sm)
	//s:=strconv.Itoa(ss)
	fmt.Printf("ffmpeg -ss %02d:%02d:%02d -t %02d:%02d:%02d",hint,mint,sint, kh, km, ks)

}
func toSec(hh, mm, ss string) int {
	var sec int
	h, _ := strconv.Atoi(hh)
	m, _ := strconv.Atoi(mm)
	s, _ := strconv.Atoi(ss)
	sec = h*3600 + m*60 + s
	return sec
}
func secTo(i int) (h, m, s int) {
	//i=3661
	h = i / 3600
	//h=1
	m = (i - h*3600) / 60
	s = (i - (h * 3600)) % 60
	return h, m, s
}
