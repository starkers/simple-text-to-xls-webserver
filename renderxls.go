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
	f.SetCellValue(sheetName, "D1", "Notes")
	offSet := 2
	for index, data := range input {
		a := fmt.Sprintf("A%d", index+offSet)
		b := fmt.Sprintf("B%d", index+offSet)
		c := fmt.Sprintf("C%d", index+offSet)
		f.SetCellValue(sheetName, a, data.Person)
		f.SetCellValue(sheetName, b, data.Time)
		f.SetCellValue(sheetName, c, data.Text)
	}

	names := getPersonNames(input)
	fmt.Println(names)
	colors := []string{
		"cfe3ff",
		"e4ffe0",
		"ffe8e0",
		"feffe0",
	}

	style0, _ := f.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{colors[0]}}})
	style1, _ := f.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{colors[1]}}})
	// style2, _ := f.NewStyle(&excelize.Style{Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{colors[2]}}})
	// Font: &excelize.Font{Color: "666666"},

	if len(names) > 0 {
		for index, data := range input {
			if data.Person == names[0] {
				f.SetRowStyle(sheetName, index+offSet, index+offSet, style0)
			}
			if data.Person == names[1] {
				f.SetRowStyle(sheetName, index+offSet, index+offSet, style1)
			}
		}
	}

	// f.SetRowStyle(sheetName, 3, 3, style1)
	err := f.SaveAs(outputFile)
	return err
}

func getPersonNames(input []Line) (result []string) {
	for _, i := range input {
		if !contains(result, i.Person) {
			result = append(result, i.Person)
		}
	}
	fmt.Println(result)
	return result
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
