package util

import "testing"

func TestGetFiles(t *testing.T) {
	dir := conf.GetValue("target", "src")
	ret := GetFiles(dir)
	t.Log(ret)
}
func TestMakeList(t *testing.T) {
	dir := conf.GetValue("target", "src")
	ret := GetFiles(dir)
	MakeList(ret)
}
