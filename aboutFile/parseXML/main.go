package main

import (
	"fmt"
	"github.com/beevik/etree"
	"log"
	"os"
	"os/exec"
	"strings"
	//"strings"
)

func init() {

}
func main() {

	xmlFile := "node.xml"
	UseEtree(xmlFile)
}
func UseEtree(fname string) {
	var Error *log.Logger
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	defer func() {
		if err := recover(); err != nil {
			Error.Println(err)
		}
	}()
	// 初始化根节点
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(fname); err != nil {
		panic(err)
	}
	root := doc.SelectElement("rss")
	servers := root.SelectElement("channel")
	for _, server := range servers.SelectElements("item") {
		//fmt.Println(server.SelectElement("title").Text())
		name := server.SelectElement("title").Text()
		//ret:=GetName(name)
		//fmt.Printf("ret = %v\n",ret)
		url := server.SelectElement("guid").Text()
		fmt.Printf("url = %v\n", url)
		cmd := exec.Command("wget", "-O", name+".m4a", url)
		if err := cmd.Run(); err != nil {
			fmt.Printf("运行命令时产生错误 %v", err)
		}
	}

}
func GetName(s string) string {
	if len(strings.Split(s, "《")) <= 1 {
		return ""
	}
	prefix := strings.Split(s, "《")[1]
	suffix := strings.Split(prefix, "》")[0]
	return suffix
	//	//tracer := "《败家子儿》 郭德纲于谦"
	//	prefix:=strings.Split(s,"《")
	//	//var (
	//	//	start int
	//	//	end int
	//	//)
	//	//
	//	//toRune := []rune(s)
	//	//for i,v:=range toRune{
	//	//	if string(v)=="《"{
	//	//		start =i
	//	//	}
	//	//	if string(v)=="》"{
	//	//		end=i
	//	//	}
	//	//
	//	//}
	//	//ret:=string([]rune(s)[start+1:end])
	//	return ret
}
