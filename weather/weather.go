package weather

import (
	"Widget/util/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

var (
	key = "3e469a45a72b31c120b0baf6609ae877"
)

type location struct {
	Status    string `json:"status"`
	Count     string `json:"count"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Forecasts []struct {
		City       string `json:"city"`
		Adcode     string `json:"adcode"`
		Province   string `json:"province"`
		Reporttime string `json:"reporttime"`
		Casts      []struct {
			Date         string `json:"date"`
			Week         string `json:"week"`
			Dayweather   string `json:"dayweather"`
			Nightweather string `json:"nightweather"`
			Daytemp      string `json:"daytemp"`
			Nighttemp    string `json:"nighttemp"`
			Daywind      string `json:"daywind"`
			Nightwind    string `json:"nightwind"`
			Daypower     string `json:"daypower"`
			Nightpower   string `json:"nightpower"`
		} `json:"casts"`
	} `json:"forecasts"`
}

var week = map[string]string{
	"1": "星期一",
	"2": "星期二",
	"3": "星期三",
	"4": "星期四",
	"5": "星期五",
	"6": "星期六",
	"7": "星期日",
}

const (
	北京市石景山区     = "110107"
	黑龙江省大庆市让胡路区 = "230604"
	海南省琼海市      = "469002"
)

func Weather() {
	m := make(map[string]string)
	//m["全国"] = "100000"
	m["北京市石景山区"] = "110107"
	m["黑龙江省大庆市让胡路区"] = "230604"
	m["海南省琼海市"] = "469002"
	for k, v := range m {
		log.Info.Printf("%s的天气信息如下\n", k)
		format(v)
	}
}
func format(adcode string) {
	//var m map[string]interface{}
	var l location
	url := "https://restapi.amap.com/v3/weather/weatherInfo?output=json&extensions=all&city=" + adcode + "&key=" + key
	response, _ := http.Get(url)
	defer response.Body.Close() //在回复后必须关闭回复的主体
	body, _ := ioutil.ReadAll(response.Body)

	if err := json.Unmarshal(body, &l); err != nil {
		log.Warn.Panicln("json解析出错")
	}
	//dump.P(l)
	log.Info.Printf("%s发布预报\n", l.Forecasts[0].Reporttime)
	for i, cast := range l.Forecasts[0].Casts {
		var witchDay string
		switch i {
		case 0:
			witchDay = "今天"
		case 1:
			witchDay = "明天"
		case 2:
			witchDay = "后天"
		case 3:
			witchDay = "大后天"
		}
		log.Info.Printf("%s%s\n", cast.Date, week[cast.Week])
		log.Info.Printf("%s白天%s%s\u2103,%s风%s级\n", witchDay, cast.Dayweather, cast.Daytemp, cast.Daywind, cast.Daypower)
		log.Info.Printf("%s夜晚%s%s\u2103,%s风%s级\n", witchDay, cast.Nightweather, cast.Nighttemp, cast.Nightwind, cast.Nightpower)
	}
	fmt.Println()
}

const (
	bajiaozhongli = "Beijing+Shijingshan+bajiao"
	qihoo         = "Beijing+Chaoyang+Wangjing"
)

// wget wttr.in/Beijing+Chaoyang+Jiuxianqiao.png
func WeatherPNG() {
	locations := []string{bajiaozhongli, qihoo}
	for _, location := range locations {
		cmd(location)
	}

}
func cmd(s string) {
	defer log.Info.Println("生成图片进程结束")
	perfix := strings.Join([]string{"wttr.in", s}, "/")
	suffix := strings.Join([]string{perfix, "png"}, ".")
	cmd := exec.Command("wget", suffix)
	log.Info.Printf("生成的命令是:%s", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Info.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		log.Info.Printf("cmd.Run产生的错误:%v", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.CMD.Println(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.CMD.Println("命令执行中有错误产生", err)
	}

}
