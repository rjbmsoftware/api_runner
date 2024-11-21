package main

import "testing"

func TestRunTest(t *testing.T) {
	url := "http://example.com"
	pathParameters := make(map[string]string)
	test := Test{
		PathParameters:       pathParameters,
		ExpectedResponseCode: 200,
	}
	headers := make(map[string]string)
	details := TestRunDefinitionDetails{
		url,
		RestMethod("GET"),
		headers,
	}
	run_test(&test, details)
}
