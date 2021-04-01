package util

import "testing"

func TestFlist(t *testing.T) {
	ret := GetFiles()
	t.Logf("ret is %v", ret)
}
