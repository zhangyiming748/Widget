package getFileType

import (
	"testing"
)

func TestDetect(t *testing.T) {
	var fp string = "../forWin32.exe"
	p := Detect(fp)
	t.Log(p)
}
