package bs4

import (
	"Widget/util/file"
	"Widget/util/log"
	"io/ioutil"
	"strings"
)

func Replace(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "《", "", -1)
	str = strings.Replace(str, "》", "", -1)
	str = strings.Replace(str, "【", "", -1)
	str = strings.Replace(str, "】", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	return str
}
func SaveHtml(s string) {
	content := []byte(s)
	err := ioutil.WriteFile("temporary.html", content, 0766)
	if err != nil {
		log.Debug.Panicf("写html文件发生错误:%s\n", err.Error())
	}
}
func ReadTemporary() string {
	filepath := "temporary.html"
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Debug.Panicf("读html缓存文件发生错误:%s\n", err.Error())
	}
	return string(content)
}
func Md2Links(src string) {
	lines := file.ReadLink(src)
	var links []string
	for _, line := range lines {
		if !strings.Contains(line, "|") {
			continue
		}
		values := strings.Split(line, "|")
		for _, value := range values {
			if strings.Contains(value, "https://www.pornhub.com/view_video.php") {
				links = append(links, value)
			}
		}
	}
	log.Info.Printf("提取出的下载链接是: %v", links)

	file.WriteLines("out.txt", links)
}
