package main

import (
	"fmt"
	"log"
	"os"
)

func readTestRunDefinition(filePath string) (*TestRunDefinition, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println("here I am")
	fmt.Println((file))
	// read the file
	// validate it is JSON
	// return what is needs?
	return nil, nil
}
