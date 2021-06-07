package util

import "testing"

func TestGetFiles(t *testing.T) {
	dir := "/Users/zen/Downloads/Downie"
	pattern := "xm"
	ret := GetFiles(dir, pattern)
	t.Log(ret)
}
