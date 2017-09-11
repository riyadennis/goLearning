package main

import (
	"fmt"
	"bufio"
	"os"
)

const fileName = "typeInfo.txt"

func main() {
	file, size := openExistingFile()
	defer file.Close()
	if size > 0 {
		writeToFile(file, os.Stdin)
	} else {
		var name string
		fmt.Print("Please enter your name: ")
		fmt.Scanf("%s", &name)
		fmt.Printf("Hello %s \xF0\x9F\x98\x83 \n", name)
		newFile, err := os.Create(fileName)
		newFile.WriteString(name + " is writing this file \n")
		if err != nil {
			fmt.Print("Unable to create the file")
		}
		writeToFile(*newFile, os.Stdin)
	}
}
func openExistingFile() (os.File, int64) {
	fmt.Println("Please type in the terminal  \xF0\x9F\x98\x83 \n")
	fileName := fileName
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return os.File{}, 0
	}
	fInfo, err := file.Stat()
	size := fInfo.Size()
	return *file, size
}
func writeToFile(file os.File, input *os.File) {
	s := bufio.NewScanner(input)
	for s.Scan() {
		if err := s.Err(); err != nil {
			fmt.Fprint(os.Stderr, "Reading error from input", err)
		} else {
			file.WriteString(s.Text() + "\n")
		}
	}
}
