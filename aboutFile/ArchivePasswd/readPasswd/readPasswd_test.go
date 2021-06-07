package readPasswd

import "testing"

func TestReadPasswd2Slicek(t *testing.T) {
	ret := ReadPasswd2Slice("./pass.txt")
	t.Log(ret)
}
