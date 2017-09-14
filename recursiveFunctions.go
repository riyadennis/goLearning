package main

import "fmt"
//recursive function
func factorial(n int)int{
	if n == 0 {
		return 1
	}else {
		return factorial(n-1) * n
	}
}

//iterative function
func square(n int ) int {
	if n == 1 {
		return 1
	} else {
		return n * n
	}
}

func main(){
	fmt.Println(factorial(5))
	fmt.Println(square(3))
}
