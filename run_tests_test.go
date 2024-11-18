package main

import "testing"

func TestRunTest(t *testing.T) {
	url := "http://example.com"
	pathParameters := make(map[string]string)
	test := Test{
		PathParameters:       pathParameters,
		ExpectedResponseCode: 200,
	}

	run_test(&test, url, RestMethod("GET"))
}
