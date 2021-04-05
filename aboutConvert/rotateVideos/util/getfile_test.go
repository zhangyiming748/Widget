package util

import "testing"

func TestGetFiles(t *testing.T) {
	dir:="/Users/zen/Downloads/Downie"
	pattern:="mp4"
	ret:=GetFiles(dir,pattern)
	t.Log(ret)
}

