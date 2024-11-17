package main

type RestMethod string

const (
	GET    RestMethod = "GET"
	POST   RestMethod = "POST"
	PUT    RestMethod = "PUT"
	DELETE RestMethod = "DELETE"
)

type Test struct {
	PathParameters       map[string]string `json:"pathParameters"`
	ExpectedResponseCode int               `json:"expectedResponseCode"`
}

type TestRunDefinition struct {
	Url        string     `json:"url"`
	RestMethod RestMethod `json:"restMethod"`
	Tests      []Test     `json:"tests"`
}
