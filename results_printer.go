package main

import (
	"errors"
	"fmt"
	"strings"
)

var fileHeader = "<html>\n<body>\n"
var fileFooter = "</body>\n</html>"

func ResultsPrinter(results *[]TestResult) (string, error) {
	if results == nil || len(*results) == 0 {
		return "", errors.New("no tests to print")
	}

	var output strings.Builder
	output.WriteString(fileHeader)

	for index, test := range *results {
		output.WriteString("<div>\n")
		output.WriteString(fmt.Sprintf("<h2>Test %d</h2>\n", index))
		output.WriteString(fmt.Sprintf("URL: %s</br>", test.Url))
		output.WriteString(fmt.Sprintf("Method %s</br>\n", test.RestMethod))

		output.WriteString("Request headers\n")
		output.WriteString("<ul>\n")
		for headerKey, headerValue := range test.RequestHeaders {
			output.WriteString(fmt.Sprintf("<li>%s : %s</li>\n", headerKey, headerValue))
		}
		output.WriteString("</ul>\n")

		output.WriteString(fmt.Sprintf("Request body <p>%s</p>\n", test.RequestBody))
		output.WriteString(fmt.Sprintf("Response code expected: %d</br>\n", test.ResponseCodeExpected))
		output.WriteString(fmt.Sprintf("Response code actual: %d</br>\n", test.ResponseCodeActual))

		output.WriteString("Response headers\n")
		output.WriteString("<ul>\n")
		for headerKey, headerValue := range test.ResponseHeaders {
			output.WriteString(fmt.Sprintf("<li>%s : %s</li>\n", headerKey, headerValue))
		}
		output.WriteString("</ul>\n")

		output.WriteString(fmt.Sprintf("Response body <p>%s</p>\n", test.ResponseBody))
		output.WriteString("</div>\n")
	}

	output.WriteString(fileFooter)

	return output.String(), nil
}
