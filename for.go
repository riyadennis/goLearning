package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for i := 7; i < 10; i++ {
		fmt.Println(i)
	}
	for {
		fmt.Println("loop")
		break
	}
	copyFile("test.txt", "destination.txt")
}
func copyFile(sourceName, destinationName string) (written int64, err error) {
	src, err := os.Open(sourceName)
	if err != nil {
		return
	}
	defer src.Close()
	dest, err := os.Create(destinationName)
	if err != nil {
		return
	}
	written, err = io.Copy(dest, src)
	defer dest.Close()
	return
}
