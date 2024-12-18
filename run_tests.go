package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func RunTestDefinition(trd *TestRunDefinition) ([]TestResult, []error) {
	details := TestRunDefinitionDetails{
		trd.Url,
		trd.RestMethod,
		trd.RequestHeaders,
	}

	results := []TestResult{}
	errors := []error{}

	for _, test := range trd.Tests {
		result, test_error := run_test(&test, details)

		if result != nil {
			results = append(results, *result)
		}

		if test_error != nil {
			errors = append(errors, test_error)
		}
	}

	return results, errors
}

func requestBody(body any) (string, io.Reader, error) {
	requestBody, err := json.MarshalIndent(body, "", "\t")
	if err != nil {
		return "", nil, err
	}

	requestBodyString := string(requestBody)
	requestBodyReader := bytes.NewReader(requestBody)

	return requestBodyString, requestBodyReader, nil
}

func run_test(test *Test, details TestRunDefinitionDetails) (*TestResult, error) {
	client := &http.Client{}
	url := details.Url

	for key, value := range test.PathParameters {
		url = strings.Replace(url, key, value, -1)
	}

	var err error
	var requestBodyString string
	var requestBodyReader io.Reader
	if test.RequestBody != nil {
		requestBodyString, requestBodyReader, err = requestBody(test.RequestBody)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(string(details.RestMethod), url, requestBodyReader)
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

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseBodyString string
	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		var out bytes.Buffer
		err = json.Indent(&out, responseBodyBytes, "", "\t")
		if err != nil {
			return nil, err
		}
		responseBodyString = out.String()
	} else {
		responseBodyString = string(responseBodyBytes)
	}

	result := &TestResult{
		test.Name,
		url,
		details.RestMethod,
		details.RequestHeaders,
		requestBodyString,
		resp.StatusCode,
		resp.Header,
		responseBodyString,
	}

	return result, nil
}
