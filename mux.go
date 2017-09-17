package main

import (
	"sync/atomic"
	"fmt"
	"time"
)

func main() {
	var o uint64 = 0
	go func() {
		for {
			for i := 0; i < 50; i++ {
				atomic.AddUint64(&o, 1)
				time.Sleep(time.Millisecond)
			}
		}
	}()

	time.Sleep(time.Second * 2)

	aFinal := atomic.LoadUint64(&o)
	fmt.Println(aFinal)
}
