package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	go generate(c)
	for {
		i := <-c
		fmt.Println(i)
		oc := make(chan int, 1)
		go filter(c, oc, i)
		c = oc
	}
}
func filter(input, output chan int, prime int) {
	for {
		i := <-input
		if i%prime != 0 {
			output <- i
		}
	}
}
func generate(c chan int) {
	for i := 2; ; i++ {
		c <- i
	}
}
