package main

type RestMethod string

const (
	GET    RestMethod = "GET"
	POST   RestMethod = "POST"
	PUT    RestMethod = "PUT"
	DELETE RestMethod = "DELETE"
)

type Test struct {
	pathParameters       map[string]string `json:"pathParameters`
	expectedResponseCode int               `json:"expectedResponseCode"`
}

type TestRunDefinition struct {
	url        string     `json:"url"`
	restMethod RestMethod `json:"restMethod"`
	tests      Test       `json:"tests"`
}
