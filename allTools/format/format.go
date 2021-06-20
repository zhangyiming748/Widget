package format

import (
	"allTools/util/log"
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func FormatSymbol(src string) {

	b := readOld(src)
	if nb := writeNew(src, b); nb != nil {
		log.Info.Println("写文件有错误")
	}

}
func readOld(src string) []byte {
	b, e := ioutil.ReadFile(src)
	if e != nil {
		log.Info.Println("read file error")
		return []byte{}
	}
	log.Info.Printf("读取到“”‘’的全部文字%s", string(b))
	sb := string(b)
	strings.Replace(sb, "“", "\"", -1)

	log.Info.Printf("替换后的全部文字%s", string(b))
	return b
}
func writeNew(src string, outPut []byte) error {
	f, err := os.OpenFile(src, os.O_WRONLY|os.O_TRUNC, 0600)

	if err != nil {
		return err
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
