package mylog

import "testing"

func TestLogof2(t *testing.T) {
	Debug.Println("我是一条信息")
	Error.Println("我是一条错误")
}
func BenchmarkDebug(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Debug.Printf("我是第%d条信息", i)
	}
	b.StopTimer()
}
func BenchmarkError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Error.Printf("我是第%d条错误", i)
	}
	b.StopTimer()
}
