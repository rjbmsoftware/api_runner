package main

import (
	"log"
)

func main() {
	testDefinition, err := readTestRunDefinition("./test.json")
	if err != nil {
		log.Println(err)
		return
	}

	run_test_definition(testDefinition)
}
