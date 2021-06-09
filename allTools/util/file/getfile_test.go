package util

import "testing"

func TestGetFiles(t *testing.T) {
	dir := "/Users/zen/Downloads/Downie"
	pattern := "xm"
	ret := GetFiles(dir, pattern)
	t.Log(ret)
}
func TestReadLine(t *testing.T) {
	src := "/Users/zen/Downloads/passwd.txt"
	ret := ReadLine(src)
	t.Log(ret)
}
