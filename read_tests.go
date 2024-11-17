package main

import (
	"encoding/json"
	"log"
	"os"
)

func readTestRunDefinition(filePath string) (*TestRunDefinition, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var testRunDefinition TestRunDefinition
	err = json.Unmarshal(file, &testRunDefinition)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &testRunDefinition, err
}
