package soup

import (
	"Widget/util/file"
	"Widget/util/log"
	"os"
	"strings"
	"testing"
)

func TestUOS(t *testing.T) {
	var urls = []string{"https://cdimage-download.chinauos.com",
		"https://cdimage-download.chinauos.com/education",
		"https://cdimage-download.chinauos.com/home",
		"https://cdimage-download.chinauos.com/home-ditch",
		"https://cdimage-download.chinauos.com/professional-wayland",
		"https://cdimage-download.chinauos.com/sp1-fix"}
	for _, url := range urls {
		UOS(url)
	}
}

func UOS(url string) {
	links := make([]string, 0)
	links = append(links, "|iso|")
	links = append(links, "|:---:|")
	resp, err := Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := HTMLParse(resp)
	isos := doc.FindAll("a")
	for _, iso := range isos {
		href := iso.Attrs()["href"]
		link := strings.Join([]string{url, href}, "/")
		if strings.Contains(link, "iso") {
			log.Info.Println(link)
			line := strings.Join([]string{"|[", href, "](", link, ")|"}, "")
			links = append(links, line)
		} else {
			log.Debug.Printf("link %v\n", link)
		}
	}
	file.WriteLines("UOS.md", links)
}
