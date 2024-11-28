package main

import (
	"strings"
	"testing"
)

func TestRunTest(t *testing.T) {
	url := "http://example.com"
	pathParameters := make(map[string]string)
	test := Test{
		PathParameters: pathParameters,
	}
	headers := make(map[string]string)
	details := TestRunDefinitionDetails{
		url,
		RestMethod("GET"),
		headers,
	}
	run_test(&test, details)
}

func TestRequestBodyValid(t *testing.T) {
	body := `{"key":"value"}`
	bodyString, bodyReader, err := requestBody(body)

	if err != nil {
		t.Error("Valid body should not error")
	}

	if strings.Compare(body, bodyString) == -1 {
		t.Errorf("Expected %s, actual %s", body, bodyString)
	}

	if bodyReader == nil {
		t.Error("Body reader should be a io.Reader not nill")
	}
}
