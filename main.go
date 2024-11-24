package main

import (
	"flag"
	"log"
)

func main() {
	filePathMessage := "f sets the path to the test definition json file"
	defaultTestFilePath := "./test.json"
	testFilePath := flag.String("f", defaultTestFilePath, filePathMessage)
	flag.Parse()

	testDefinition, err := readTestRunDefinition(*testFilePath)
	if err != nil {
		log.Println(err)
		return
	}

	results, errors := RunTestDefinition(testDefinition)
	if len(errors) > 0 {
		for _, testErrors := range errors {
			log.Println(testErrors)
		}
	}

	err = ResultsPrinter(&results)
	if err != nil {
		log.Println(err)
	}
}
