package delMacos

import (
	util "allTools/util/conf"
	. "allTools/util/file"
	"testing"
)

func TestGetFiles(t *testing.T) {
	src := "/Volume/MI"
	p := util.GetVal("location", "pattern")
	GetFiles(src, p)
}
func TestDelFile(t *testing.T) {
	delFiles("/Volume/MI")
}
