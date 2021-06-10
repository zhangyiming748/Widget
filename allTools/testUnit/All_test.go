package testUnit

import (
	"fmt"
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
