package main

import (
	"errors"
)

var fileHeader = "<html>"
var fileFooter = "</html>"

func ResultsPrinter(results *[]TestResult) (string, error) {

	if results == nil || len(*results) == 0 {
		return "", errors.New("no tests to print")
	}

	return fileHeader + fileFooter, nil
}
