package timeafter

import (
	"log"
	"time"
)

func subTime() {
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()
	sub := t2.Sub(t1)
	log.Printf("sub = %v\tsub`s type is %T", sub, sub)
}
