package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(time.Second * 2)

		c2 <- "Two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Printf("Recived %s\n", msg)
		case msg1 := <-c2:
			fmt.Printf("Recieved %s \n", msg1)
		}
	}
}
