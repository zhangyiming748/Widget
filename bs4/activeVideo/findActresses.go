package activeVideo

import (
	"Widget/bs4"
	"Widget/bs4/soup"
	"Widget/util/file"
	"Widget/util/log"
	"strings"
)

type Actresses struct {
	Name  string
	Alias string
	Image string
}

func FindActresses(url string) {
	lines := make([]string, 0)
	imgs := make([]string, 0)
	resp, err := soup.GetWithProxy(url, "http://127.0.0.1:8889")
	if err != nil {
		log.Info.Panicf("err = %v\n", err)
	}
	bs4.SaveHtml(resp)
	doc := soup.HTMLParse(resp)
	boxes := doc.Find("div", "class", "actors")
	log.Info.Println(boxes)
	divs := boxes.FindAll("div")
	log.Info.Println(divs)
	for i, div := range divs {
		a := div.Find("a")
		alias := a.Attrs()["title"]
		alias = bs4.Replace(alias)

		figure := a.Find("figure", "class", "image")
		img := figure.Find("img", "class", "avatar")
		src := img.Attrs()["src"]

		strong := a.Find("strong")
		name := strong.Text()
		name = bs4.Replace(name)

		log.Info.Printf("第%d位老师%s的别名是%s\t对应的头像是%s\n", i+1, name, alias, src)
		var act = &Actresses{
			Name:  name,
			Alias: alias,
			Image: src,
		}
		fname := strings.Join([]string{act.Name, "jpg"}, ".")
		if strings.Contains(src, "unknow") {
			fname = strings.Join([]string{act.Name, "png"}, ".")
			act.Image = "https://javdb.com/assets/actor_unknow-15f7d779b3d93db42c62be9460b45b79e51f8a944796eee30ed87bbb04de0a37.png"
		}
		log.Info.Printf("文件名:%s\n", fname)
		downlink := strings.Join([]string{"wget", act.Image, "-O", fname}, " ")
		imgs = append(imgs, downlink)
		line := strings.Join([]string{"|", act.Name, "|", act.Alias, "|", "![", act.Name, "](", act.Image, ")|"}, "")
		lines = append(lines, line)
	}
	file.WriteLines("actresses.sh", imgs)
	file.WriteLines("actresses.md", lines)
}
