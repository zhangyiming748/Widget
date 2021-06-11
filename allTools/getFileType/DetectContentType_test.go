package getFileType

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"os"
	"testing"
)

func TestDetect(t *testing.T) {
	var fp string = "ansi.txt"
	p := Detect(fp)
	t.Log(p)
}
func TestRead(t *testing.T) {
	fi, err := os.Open("ansi.txt")
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
}
func TestGBK2UTF8(t *testing.T) {
	GBK2UTF8("ansi.txt", "utf8.txt")
}
