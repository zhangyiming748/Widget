package unzip

import "testing"

func TestUnzip(t *testing.T) {
	src := "/Users/zen/Desktop/EverymacOScombined.zip"
	dst := "/Users/zen/Downloads/Downie/"
	pwd := "password"
	UnZip(src, dst, pwd)
}
func TestDelFail(t *testing.T) {
	dst := "/Users/zen/Downloads/Downie"
	delFail(dst)
}
