package downloadcmd

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"sync"
	"time"
	"youtube-dl/convert"
	. "youtube-dl/mylog"
	"youtube-dl/util"
)

func RunCmd(url string, wg *sync.WaitGroup, proxy, dir string, i int, ch chan struct{}) {
	path := strings.Join([]string{dir, "%(title)s.%(ext)s"}, "/")
	cmd := exec.Command("youtube-dl", "--proxy", proxy, "-o", path, "-f", "best", url)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		Error.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		Error.Printf("cmd.Run产生的错误:%v", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if strings.HasSuffix(string(tmp),"has already been downloaded"){
			Debug.Println("当前文件已存在")
		}
		Info.Printf("第%d个文件输出:%s", i+1, string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		Error.Printf("%v对应文件:%v\n", err, url)
	}
	Info.Printf("下载文件%v完成\n", url)
	json := getJson(url, proxy)
	title := parseJson(json)
	f, notfound := getCurrentFile(title)
	if notfound != nil {
		Error.Println("待转换的文件未找到")
	}
	fp := getFullpath(f)
	if isM3u8(fp){
		Debug.Printf("准备转换文件:%s\n",f)
		convert.Convert(fp)
		Debug.Printf("转换文件%s完成",f)
	}
	//Debug.Printf("fp = %s\n",fp)
	wait := time.Duration(rand.Intn(3))
	time.Sleep(wait * time.Second)
	<-ch
	wg.Done()
}

//获取文件的json格式详情
func getJson(url, proxy string) []byte {
	//path := strings.Join([]string{dir, "%(title)s.%(ext)s"}, "/")
	cmd := exec.Command("youtube-dl", "--proxy", proxy, "-j", url)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		Error.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		Error.Printf("cmd.Run产生的错误:%v", err)
	}
	var static []byte
	var buffer bytes.Buffer
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if err != nil {
			break
		}
		buffer.Write(static)
		buffer.Write(tmp)
		static = buffer.Bytes() //得到了b1+b2的结果
	}
	return static
}

//解析json并返回不带扩展名的文件名
func parseJson(b []byte) string {
	sj := string(b)
	title := gjson.Get(sj, "title").String()
	return title
}

//遍历当前文件夹找到正在运行下载的文件
func getCurrentFile(name string) (string, error) {
	workdir := util.GetVal("target", "fp")
	fileInfoList, err := ioutil.ReadDir(workdir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		fname := fileInfoList[i].Name()
		if strings.HasPrefix(fname, name) {
			return fname, nil
		}
	}
	return "no value", err
}

//判断是否为M3U8文件
func isM3u8(file string) bool {
	if ret := strings.Split(file, "."); ret[len(ret)-1] == "m3u8" {
		return true
	} else {
		return false
	}
}

//拼接绝对路径
func getFullpath(f string) string {
	fp := strings.Join([]string{util.GetVal("target", "fp"), f}, "/")
	return fp
}
