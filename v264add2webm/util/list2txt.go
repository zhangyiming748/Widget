//package util
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//func doInit(f string) {
//	//f:=conf.GetValue("target","list")
//
//	os.Create(f)
//}
//func MakeList(fs []string) {
//	var f string = GetVal("target", "list")
//
//	for i, v := range fs {
//		writeLine(v, f)
//		fmt.Printf("写入第%d行\n", i)
//	}
//}
//func writeLine(s, file string) {
//	var str = strings.Join([]string{s, "\n"}, "")
//	var filename = file
//	var f *os.File
//	var err1 error
//	if f, err1 = os.OpenFile(filename, os.O_APPEND, 0666); err1 != nil {
//		fmt.Println(err1)
//	}
//
//	defer f.Close()
//	w := bufio.NewWriter(f) //创建新的 Writer 对象
//	n, err := w.WriteString(str)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	fmt.Printf("写入 %d 个字节\n", n)
//	w.Flush()
//}

package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MakeList(fs []string) {
	var f string = GetVal("target", "list")

	writeLine(fs, f)

}
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
func writeLine(s []string, file string) {

	var f *os.File
	var err1 error
	if checkFileIsExist(file) { //如果文件存在

		f, err1 = os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(file) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	if err1 != nil {
		panic(err1)
	}
	for _, v := range s {
		w := bufio.NewWriter(f) //创建新的 Writer 对象
		w.WriteString("file ")
		fv := strings.Join([]string{GetVal("target", "src"), v}, "/")
		n, _ := w.WriteString(fv)
		w.WriteString("\n")
		fmt.Printf("写入 %d 个字节n", n)
		w.Flush()
	}
}
