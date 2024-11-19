package main

import (
	"log"
	"net/http"
)

func run_test_definition(test_run_definition *TestRunDefinition) {
	url := test_run_definition.Url
	rest_method := test_run_definition.RestMethod
	tests := test_run_definition.Tests

	for _, test := range tests {
		run_test(&test, url, rest_method)
	}
}

func run_test(test *Test, url string, rest_method RestMethod) {
	client := &http.Client{}

	req, err := http.NewRequest(string(rest_method), url, nil)
	if err != nil {
		log.Printf("things broken")
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("more things are broken")
	}

	log.Println(url)
	log.Printf("actual response code: %d, expected: %d", resp.StatusCode, test.ExpectedResponseCode)
}
