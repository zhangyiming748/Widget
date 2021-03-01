package util

import "testing"

func TestGetVal(t *testing.T) {
	ret := GetVal("proxy", "port")
	t.Logf("获取到的配置信息:%v\n", ret)
}
