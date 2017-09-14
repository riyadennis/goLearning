package main

import "fmt"

func zeroval(i int) {
	fmt.Println("value wont change")
	 i = i * 5
}
func zeroptr(i *int) {
	fmt.Println("value will change")
	  *i = *i * 5
}
func main() {
	i := 1;
	zeroval(i)
	fmt.Println(i)
	zeroptr(&i)
	fmt.Println(i)
	fmt.Println(&i)
}
