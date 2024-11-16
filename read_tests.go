package main

import (
	"fmt"
	"log"
	"os"
)

func readTests(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(file)
	// read the file
	// validate it is JSON
	// return what is needs?
}
