package util

import (
	"allTools/util/log"
	"bufio"
	"io"
	"os"
)

func ReadLine(src string) []string {
	fi, err := os.Open(src)
	if err != nil {
		log.Info.Printf("打开字典文件失败: %s\n", err)
		return []string{}
	}
	defer func() {
		if err := fi.Close();err != nil {
			log.Info.Printf("关闭字典文件失败: %s\n", err)
		}
	}()
	pswds := []string{}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		pswds = append(pswds, string(a))
		log.Info.Printf("读取到的密码(%s)\n", string(a))
	}
	return pswds
}
