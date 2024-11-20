package main

import (
	"log"
	"net/http"
	"strings"
)

func run_test_definition(test_run_definition *TestRunDefinition) {
	url := test_run_definition.Url
	rest_method := test_run_definition.RestMethod
	tests := test_run_definition.Tests
	headers := test_run_definition.RequestHeaders

	for _, test := range tests {
		run_test(&test, url, rest_method, headers)
	}
}

func run_test(test *Test, url string, rest_method RestMethod, headers map[string]string) {
	client := &http.Client{}

	for key, value := range test.PathParameters {
		url = strings.Replace(url, key, value, -1)
	}
	req, err := http.NewRequest(string(rest_method), url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for header, value := range headers {
		req.Header.Add(header, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(test.Name)
	log.Println(url)
	log.Printf("actual response code: %d, expected: %d", resp.StatusCode, test.ExpectedResponseCode)
	log.Println(resp.Body)
	log.Println("Response headers")
	for header, value := range resp.Header {
		log.Println(header, value)
	}
	log.Print("\n\n\n")
}
