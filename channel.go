package main

import (
	"os"
	"fmt"
)

type lookup struct{
	name string
	isCloudFlare bool
}

func main () {
	// open files r and w
	file, err := os.OpenFile("typeInfo2017-Sep-11.txt", os.O_APPEND|os.O_WRONLY,0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(" this is the APPENDED text"); err != nil {
		panic(err)
	}

	fmt.Printf("Appended into file\n")
}