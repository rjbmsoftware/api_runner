package main

import (
	"io"
	"net/http"
	"strings"
)

func run_test_definition(trd *TestRunDefinition) {
	details := TestRunDefinitionDetails{
		trd.Url,
		trd.RestMethod,
		trd.RequestHeaders,
	}

	for _, test := range trd.Tests {
		run_test(&test, details)
	}
}

func run_test(test *Test, details TestRunDefinitionDetails) (*TestResult, error) {
	client := &http.Client{}
	url := details.Url

	for key, value := range test.PathParameters {
		url = strings.Replace(url, key, value, -1)
	}

	req, err := http.NewRequest(string(details.RestMethod), url, nil)
	if err != nil {
		return nil, err
	}

	for header, value := range details.RequestHeaders {
		req.Header.Add(header, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseBodyBytes, err := io.ReadAll(resp.Request.Body)
	if err != nil {
		return nil, err
	}

	result := &TestResult{
		url,
		details.RestMethod,
		details.RequestHeaders,
		"",
		test.ExpectedResponseCode,
		resp.StatusCode,
		resp.Header,
		string(responseBodyBytes),
	}

	return result, nil
}
