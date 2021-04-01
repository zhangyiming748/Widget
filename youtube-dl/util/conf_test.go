package util

import "testing"

func TestGetVal(t *testing.T) {
	ret := GetVal("goroutine", "num")
	t.Logf("获取到的配置信息:\"%v\"\n", ret)
}
