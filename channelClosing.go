package main

import "fmt"

func main() {
	jobs := make(chan string, 5)
	done := make(chan bool, 1)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Printf("Recived job from channel %s", j)
			} else {
				fmt.Println("All jobs recieved")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		jobs <- "My job \n"
	}

	close(jobs)
	fmt.Println("All jobs sent")
	<-done

}
