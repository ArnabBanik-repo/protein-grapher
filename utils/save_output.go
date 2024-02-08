package utils

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ArnabBanik-repo/protein-grapher/models"
)

func saveToExcel(paths []models.ProteinList, drug models.Protein) {
	var outputFileName string = fmt.Sprintf("output/%s.xlsx", drug)
	file := excelize.NewFile()
	for rowIndex, rowData := range paths {
		for colIndex, cellData := range rowData {
			cell := excelize.ToAlphaString(colIndex+1) + fmt.Sprint(rowIndex+1)
			file.SetCellValue("Sheet1", cell, cellData)
		}
	}

	if err := file.SaveAs(outputFileName); err != nil {
		fmt.Printf("Error saving %q file:\n%s", outputFileName, err)
		return
	}

	fmt.Printf("%v ✅\n", drug)
}

func saveToCsv(paths []models.ProteinList, drug models.Protein, subdir string) {
	var outputFileName string = fmt.Sprintf("./output/%s/%s.csv", subdir, drug)

	file, err := os.Create(outputFileName)
	if err != nil {
		fmt.Printf("Could not create %v\n", outputFileName)
		return
	}
	csvWriter := csv.NewWriter(file)

	// Look for a way to make this better. Looks way too ugly
	for _, path := range paths {
		temp := []string{}
		for _, protein := range path {
			temp = append(temp, string(protein))
		}
		err = csvWriter.Write(temp)
		if err != nil {
			fmt.Printf("Could not write to %v\n", outputFileName)
			return
		}
	}
	csvWriter.Flush()
}
