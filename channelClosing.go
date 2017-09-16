package main

import "fmt"
//once all the processing is done and information is put to the channel we can close it
// still we cab iterate through the channel in the reciever even if the channel is closed
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
