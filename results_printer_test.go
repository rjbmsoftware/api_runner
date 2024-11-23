package main

import "testing"

func TestResultPrinter(t *testing.T) {
	expectedString := ""
	actualString, testResultError := ResultsPrinter(nil)
	expectedErrorMessage := "no tests to print"

	if actualString != expectedString {
		t.Errorf("Expected %s, received: %s", actualString, expectedString)
	}

	if testResultError == nil {
		t.Errorf("Expected error, received nil")
	}

	actualErrorMessage := testResultError.Error()
	if actualErrorMessage != expectedErrorMessage {
		t.Errorf(`Expected %s, received %s`, expectedErrorMessage, actualErrorMessage)
	}
}
