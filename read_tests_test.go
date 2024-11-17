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

func TestReadFileValid(t *testing.T) {
	path := "./test.json"
	testRunDefinition, err := readTestRunDefinition(path)

	if err != nil {
		t.Errorf("cannot read valid json file: %s", err)
	}

	if testRunDefinition == nil {
		t.Errorf("test run definition is nill")
	}

	var my_thing *TestRunDefinition
	expectedType := reflect.TypeOf(my_thing)
	actualType := reflect.TypeOf(testRunDefinition)

	if actualType != expectedType {
		t.Errorf("Types do not match expected: %s, received: %s", expectedType, actualType)
	}
}
