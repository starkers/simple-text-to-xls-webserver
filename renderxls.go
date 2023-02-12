package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func renderXls(input []Line, outputFile string) error {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheetName := "Sheet1"
	f.SetColWidth(sheetName, "A", "A", 20)
	f.SetCellValue(sheetName, "A1", "Person")
	f.SetCellValue(sheetName, "B1", "Time")
	f.SetCellValue(sheetName, "C1", "Text")
	offSet := 2
	for index, data := range input {
		a := fmt.Sprintf("A%d", index+offSet)
		b := fmt.Sprintf("B%d", index+offSet)
		c := fmt.Sprintf("C%d", index+offSet)
		f.SetCellValue(sheetName, a, data.Person)
		f.SetCellValue(sheetName, b, data.Time)
		f.SetCellValue(sheetName, c, data.Text)
	}
	err := f.SaveAs(outputFile)
	return err
}
