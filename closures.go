package main

import "fmt"

func masterFunc(i int) func() int {
	return func() int {
		i +=i
		return i
	}
}

func main() {
	myFunc := masterFunc(100)
	fmt.Println(myFunc())
	fmt.Println(myFunc())

	myFunc2 := masterFunc(20)
	fmt.Println(myFunc2())
}
