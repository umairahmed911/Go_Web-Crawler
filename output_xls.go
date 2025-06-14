package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func ExportResultsToXLS(results []CrawlResult, filename string) error {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	f.SetActiveSheet(0)
	f.SetSheetRow(sheet, "A1", &[]interface{}{"Depth", "Status", "URL", "Title"})
	for i, res := range results {
		row := []interface{}{res.Depth, res.Status, res.URL, res.Title}
		cell := fmt.Sprintf("A%d", i+2)
		f.SetSheetRow(sheet, cell, &row)
	}
	return f.SaveAs(filename)
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
