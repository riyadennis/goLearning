package main

import (
	"time"
	"fmt"
)

func main() {

	requests := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			requests <- i
		}
		close(requests)
	}()

	limiter := time.Tick(time.Millisecond * 1000)

	for r := range requests {
		<-limiter
		fmt.Println("Requests ...", r)
	}

	burstyTicker := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyTicker <- time.Now()
	}
	go func(){
		for t := range time.Tick(time.Second * 2) {
			burstyTicker <- t
		}
	}()

	burstyChan := make(chan int, 20)

	for i := 0; i < 20; i++ {
		burstyChan <- i
	}
	close(burstyChan)

	for b := range burstyChan {
		<-burstyTicker
		fmt.Println("Channel has", b, time.Now())
	}
}
