package main

import (
	"fmt"
	"testing"
)

func TestStr2(t *testing.T) {
	times := "211224"
	hh := times[0:2]
	mm := times[2:4]
	ss := times[4:]
	t.Logf("%s:%s:%s\n", hh, mm, ss)
}
func TestSecTo(t *testing.T) {
	h, m, s := secTo(3662)

	fmt.Printf("%02d:%02d:%02d", h, m, s)
}
