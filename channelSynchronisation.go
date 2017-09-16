package main

import (
	"fmt"
	"time"
)

//done with a channel that can hold a bool
// once all the processing is done pass a true value into that channel
// recieve from that channel in the main function, once we recieve anything into that channel
//execution terminates

func printInLoop(count int, done chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second)
	for i := 0; i < count; i ++ {
		fmt.Println("I am printing")
	}
	fmt.Println("Done")
	done <- true
}
func ping(c chan<- string, s string) {
	c <- s
}
func pong(c <-chan string, s chan<- string) {
	m := <-c
	s <- m
}
func main() {
	ch := make(chan bool, 1)
	go printInLoop(100, ch)

	p := make(chan string)
	po := make(chan string)
	ping(p, "hello")
	pong(p, po)
	fmt.Println(<-po)

	<-ch
}
