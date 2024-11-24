package main

import (
	_ "embed"
	"html/template"
	"os"
)

//go:embed "report.html"
var reportTemplateText string

func ResultsPrinter(results *[]TestResult) error {
	reportFile, err := os.Create("./test_report_test.html")
	if err != nil {
		return err
	}

	reportTemplate, err := template.New("report").Parse(reportTemplateText)
	if err != nil {
		return err
	}

	err = reportTemplate.Execute(reportFile, *results)
	return err
}
