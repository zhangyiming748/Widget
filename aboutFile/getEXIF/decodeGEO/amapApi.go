
//https://restapi.amap.com/v3/geocode/regeo?output=json&location=116.310003,39.991957&key=3e469a45a72b31c120b0baf6609ae877&radius=1000&extensions=all
package decodeGEO

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)
func Decode(locat string){
	url:="https://restapi.amap.com/v3/geocode/regeo?output=xml&location="+locat+"&key=3e469a45a72b31c120b0baf6609ae877&radius=1000&extensions=all"
	response, err := http.Get(url)
	if err != nil{
		//...
	}
	defer response.Body.Close()//在回复后必须关闭回复的主体
	body,err := ioutil.ReadAll(response.Body)
	if err != nil{

	}
	//fmt.Println(string(body))
	result(string(body))
}
func result(text string)  {
	if Exists("result.xml"){
		os.Remove("result.xml")
	}
	f,err := os.Create("result.xml")
	defer f.Close()
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		_,err=f.Write([]byte(text))
		if err !=nil {
			fmt.Println(err.Error())
		}
	}
}
func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}