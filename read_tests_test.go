package main

import (
	"reflect"
	"testing"
)

func TestReadFileNotFound(t *testing.T) {
	path := "./some.json"
	testRunDefinition, err := readTestRunDefinition(path)

	if testRunDefinition != nil {
		t.Errorf("reading run definition is defined expected nil")
	}

	if err == nil {
		t.Errorf("reading an invalid file path did not error: %s, %s", path, err)
	}
}

func TestReadFileTypes(t *testing.T) {
	path := "./test.json"
	testRunDefinition, err := readTestRunDefinition(path)

	if err != nil {
		t.Errorf("cannot read valid json file: %s", err)
	}

	if testRunDefinition == nil {
		t.Errorf("test run definition is nill after reading from valid path: %s", path)
	}

	var my_thing *TestRunDefinition
	expectedType := reflect.TypeOf(my_thing)
	actualType := reflect.TypeOf(testRunDefinition)

	if actualType != expectedType {
		t.Errorf("Types do not match expected: %s, received: %s", expectedType, actualType)
	}
}

func TestReadValid(t *testing.T) {
	path := "./test.json"
	testRunDefinition, err := readTestRunDefinition(path)

	if err != nil {
		t.Errorf("should not error: %s", err)
	}

	if testRunDefinition == nil {
		t.Errorf("Run definition is nil")
	}

	expectedURL := "http://example.com"
	actualURL := testRunDefinition.Url

	if actualURL != expectedURL {
		t.Errorf("test data URI mismatch expected: %s, actual: %s", expectedURL, actualURL)
	}
}
