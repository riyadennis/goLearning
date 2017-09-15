package main

import (
	"errors"
	"fmt"
)

func f (n []int) (int, error){
	var s = 0
	for _, i := range n {
		if i == 43 {
			return -1, errors.New("43 not allowed")
		} else {
			s += i
		}
	}
	return s, nil
}

//by convention errors are the last return value
//and have type error as built in interface
func main(){
	nums := []int{7,43}
	if _, err := f(nums); err != nil {
		fmt.Println("Failed")
	}
	num1 := []int{2, 45}
	if _, err := f(num1); err == nil {
		fmt.Println("Passed")
	}
}