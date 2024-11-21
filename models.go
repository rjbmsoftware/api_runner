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
	Name                 string            `json:"name"`
	RequestBody          string
}

type TestRunDefinition struct {
	Url            string            `json:"url"`
	RestMethod     RestMethod        `json:"restMethod"`
	Tests          []Test            `json:"tests"`
	RequestHeaders map[string]string `json:"requestHeaders"`
}

type TestRunDefinitionDetails struct {
	Url            string
	RestMethod     RestMethod
	RequestHeaders map[string]string
}

type TestResult struct {
	Url                  string
	RestMethod           RestMethod
	RequestHeaders       map[string]string
	RequestBody          string
	ResponseCodeExpected int
	ResponseCodeActual   int
	ResponseHeaders      map[string][]string
	ResponseBody         string
}
