package downloadcmd

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestYtd(t *testing.T) {
	var wg sync.WaitGroup
	link := ""
	wg.Add(1)
	RunCommand(link, &wg, 1)
	wg.Wait()

}


func RunCommand(link string, s *sync.WaitGroup, i int) {

}

//func TestSplit(t *testing.T) {
//	link := ""
//	ret := split(link)
//	log.Println(ret)
//}
func TestJust(t *testing.T) {
	i := 1
	s := fmt.Sprintf("第%d个管道返回:\t", i)
	t.Log(s)
	ch := strings.Join([]string{"ch", strconv.Itoa(i)}, "")
	t.Log("ch=", ch)

}
func TestGetJson(t *testing.T) {
	url:="https://www.pornhub.com/view_video.php?viewkey=ph56a17d4a0d0ad"
	proxy:="127.0.0.1:8889"
	ret:=getJson(url,proxy)
	t.Logf("内容:%v\t类型%T\n",string(ret),ret)
}
func TestGetCurrentFile(t *testing.T) {
	ret, _ :=getCurrentFile("/Users/zen/Github/Widget/youtube-dl")
	t.Log(ret)
}
func TestGetJSON(t *testing.T) {
	url:="https://www.pornhub.com/view_video.php?viewkey=ph5ee70c0f48929"
	proxy:="127.0.0.1:8889"
	ret1:=getJson(url,proxy)
	t.Logf("%s",string(ret1))
	ret2:=parseJson(ret1)
	t.Logf("title : %s",ret2)
	ret3,_:=getCurrentFile(ret2)
	t.Logf("fullname : %s",ret3)
	ret4:=isM3u8(ret3)
	t.Logf("is m3u8? : %t",ret4)
}