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
