package main

import "fmt"

func fs(s string, c int) {
	for i := 0; i < c; i ++ {
		fmt.Printf("Printing %s : %d \n", s, i)
	}
}
func main() {
	fs("hello", 10)
	//buffered channel can accepts values without a reciever
	ch := make(chan []string, 2)
	go fs("go routine", 200)

	go func() {
		//create an array of string
		chStr := make([]string, 100)
		chStr1 := make([]string, 100)
		chStr2 := make([]string, 100)
		for i := 0; i < 100; i ++ {
			chStr[i] =  "Hello from channel \n"
			chStr1[i] =  "channel2 \n"
			chStr2[i] =  "channel3 \n"
		}
		//put it to the channel has size 2 still can recieve more
		ch<-chStr
		ch<-chStr1
		ch<-chStr2
	}()


	var input string
	//recieving from the channel
	fromChan := <-ch
	fmt.Println(fromChan)

	fromChan2 := <-ch
	fmt.Println(fromChan2)

	fromChan3 := <-ch
	fmt.Println(fromChan3)

	fmt.Scanln(&input)
	fmt.Println("Done")
}
