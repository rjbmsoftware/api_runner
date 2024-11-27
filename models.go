package main

type RestMethod string

const (
	GET    RestMethod = "GET"
	POST   RestMethod = "POST"
	PUT    RestMethod = "PUT"
	DELETE RestMethod = "DELETE"
)

type Test struct {
	PathParameters map[string]string `json:"pathParameters"`
	Name           string            `json:"name"`
	RequestBody    any               `json:"requestBody"`
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
	Name               string
	Url                string
	RestMethod         RestMethod
	RequestHeaders     map[string]string
	RequestBody        string
	ResponseCodeActual int
	ResponseHeaders    map[string][]string
	ResponseBody       string
}
