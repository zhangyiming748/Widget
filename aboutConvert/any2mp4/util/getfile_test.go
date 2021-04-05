package util

import "testing"

func TestGetFiles(t *testing.T) {
	dir := "/Users/zen/Downloads/Downie"

	ret := GetFiles(dir)
	t.Log(ret)
}
