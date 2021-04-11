package util

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t1:=time.Now()
	t2:=time.Now().AddDate(0,0,2).Add(3*time.Minute)
	t.Log(t2.Sub(t1))
}
