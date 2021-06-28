package testUnit

import (
	"fmt"
	"strings"

	//"github.com/go-cmd/cmd"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	printInLine()
}
func printInLine() {
	for i := 0; i < 100; i++ {
		fmt.Printf("这是第%d次输出", i)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("\r\033[k")
	}
}
func TestSpilt(t *testing.T) {
	str1 := "abc;def;ghi"
	str2 := "xyz"
	ret1 := strings.Split(str1, ";")
	t.Logf("ret1=%v\n", ret1)
	ret2 := strings.Split(str2, ";")
	t.Logf("ret2=%v\n", ret2)
}
