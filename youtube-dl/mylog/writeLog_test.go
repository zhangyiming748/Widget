package mylog

import "testing"

func TestLogof(t *testing.T) {
	var str = "测试1\n测试2\n"
	Logof(str)
}
func BenchmarkLogof(b *testing.B) {
	var str = "测试1\n测试2\n"
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		Logof(str)
	}
	b.StopTimer()
}
