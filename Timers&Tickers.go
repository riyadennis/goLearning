package main

import (
	"time"
	"fmt"
)

func main() {
	timer1 := time.NewTimer(time.Millisecond * 100)
	<-timer1.C
	fmt.Println("Timer1 expired")

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Timer ticks at ", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
