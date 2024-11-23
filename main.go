package main

import (
	"flag"
	"log"
	"os"
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

	report, reportError := ResultsPrinter(&results)
	if reportError != nil {
		log.Println(reportError)
		return
	}

	fileContent := []byte(report)
	err = os.WriteFile("somefile.html", fileContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
