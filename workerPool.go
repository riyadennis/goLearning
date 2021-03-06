package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "Processing job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished processing", j)
		results <- j * 2
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < 5; j++ {
		jobs <- j
	}
	close(jobs)
	for r := 0; r < 5; r++ {
		<-results
	}

}
