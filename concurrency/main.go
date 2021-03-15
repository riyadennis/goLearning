package main

import (
	"fmt"
	"sync"
)

func main(){
	var data int
	var memoryAccess sync.Mutex
	go func(){
		memoryAccess.Lock()
		data ++
		memoryAccess.Unlock()
	}() // critical section 1

	memoryAccess.Lock()
	if data == 0{ // critical section 2
		fmt.Print("data is empty")
	} else{
		fmt.Printf("data: %v", data)//critical section 3
	}
	memoryAccess.Unlock()
}
