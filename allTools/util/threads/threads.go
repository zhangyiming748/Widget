package threads

import (
	"runtime"
	"strconv"
)
//根据当前的系统返回适当的线程数
func Threads() string {
	var threads int = runtime.NumCPU()
	if system := runtime.GOOS; system == "darwin" {
		threads = 1
	} else {
		threads = 4
	}
	t := strconv.Itoa(threads)
	return t
}
