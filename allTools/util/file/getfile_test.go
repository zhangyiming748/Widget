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
func TestGetMultiFile(t *testing.T) {
	dir := "/Volumes/TimeMachine/Download/hiczvko#678#[3D][Idemi-iam][无修正]Patreon 2021年5月奖励/[3D][Opiumud（OP社）][无修正]Opiumud-002"
	pattern := "wmv;rm"
	ret := GetFiles(dir, pattern)
	t.Logf("%v\n", ret)
}
