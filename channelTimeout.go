package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "Hello"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	//this is the place where we implement time out after 2 s
	// seconds we print the string and stops the execution
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Riya"
	}()

	select {
	case res1 := <-c2:
		fmt.Println(res1)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 2")
	}

	signal := make(chan string, 1)
	messages := make(chan string, 1)

	messages <- "Good Morning"
	signal <-"hello"
	select {
	case message := <-messages:
		fmt.Printf("Message recieved %s\n", message)
	default:
		fmt.Println("Nothing happend")
	}

	select{
	case s :=<-signal:fmt.Printf("Signal recived %s \n ", s)
	default:
		fmt.Println("No signal recieved")
	}
}
