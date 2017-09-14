package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const a = iota
	fmt.Printf("%d\n", a)
	fmt.Printf("hello world")

	f := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}
	http.HandleFunc("/hello", f)
	log.Panic(http.ListenAndServe(":8080", nil))
}
