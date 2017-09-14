package main

import "fmt"

//variadic function
//can take any number of arguments
//even a slice
func sum(num ...int) (int) {
	s := 0
	for _, n := range num {
		s += n
	}
	return s
}
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2))
	//integer slice
	//can be passed into a variadic function with ...
	sl := []int{1,2,3,4,5}
	fmt.Println(sum(sl...))
	sl1 := []int{2,3,6,7,8}
	fmt.Println(sum(sl1...))

	sl3 := []int{34,45,56,78,90}
	fmt.Println(sum(sl3...))
}
