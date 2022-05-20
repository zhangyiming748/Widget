package file

func ReadLink(src string) []string {
	fi, err := os.Open(src)
	if err != nil {
		log.Info.Printf("打开下载链接文件失败: %s\n", err)
		return []string{}
	}
	defer func() {
		if err := fi.Close(); err != nil {
			log.Info.Printf("关闭下载链接文件失败: %s\n", err)
		}
	}()
	links := []string{}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		links = append(links, string(a))
		log.Info.Printf("读取到的下载链接(%s)\n", string(a))
	}
	return links
}
