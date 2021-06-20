package main

import (
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	s := "《败家子儿》 郭德纲于谦"
	prefix := strings.Split(s, "《")[1]
	suffix := strings.Split(prefix, "》")[0]
	t.Logf(suffix)
}
