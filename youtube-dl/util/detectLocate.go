package util

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func detectlocate() string {

}
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	re, err := filepath.Abs(file)
	if err != nil {
		logs.Error("The eacePath failed: %s\n", err.Error())
	}
	flog.Infoln("The path is ", re)
	return filepath.Abs(file)
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
	log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
