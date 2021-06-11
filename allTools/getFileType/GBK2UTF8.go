package getFileType

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"os"
)

func GBK2UTF8(src, dst string) {
	fi, err := os.Open(src)
	if err != nil {
		return
	}
	defer fi.Close()
	decoder := mahonia.NewDecoder("gb18030") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	fd, err := ioutil.ReadAll(decoder.NewReader(fi))
	if err != nil {
		return
	}
	fmt.Println("ds", string(fd))
	write(fd, dst)
}
func write(content []byte, dst string) {

	//content := []byte("测试1\n测试2\n")
	err := ioutil.WriteFile(dst, content, 0644)
	if err != nil {
		panic(err)
	}
}
