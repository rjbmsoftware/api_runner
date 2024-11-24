package main

import (
	"html/template"
	"os"
)

func ResultsPrinter(results *[]TestResult) error {
	reportTemplate, err := template.ParseFiles("report.html")
	if err != nil {
		return err
	}

	reportFile, err := os.Create("./test_report_test.html")
	if err != nil {
		return err
	}

	err = reportTemplate.Execute(reportFile, *results)
	return err
}
