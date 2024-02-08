package db

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"

	"github.com/ArnabBanik-repo/protein-grapher/models"
)

const (
	fillPlasma int = 0
	fillSource int = 1
)

var Routes = models.ProteinNetwork{}
var Plasma = models.ProteinList{}
var Sources = models.ProteinList{}

var AllProteins = models.ProteinList{}

func FillVar(lines [][]string, fillVar int) {
	switch fillVar {
	case 0:
		for i := range lines {
			Plasma = append(Plasma, models.Protein(lines[i][0]))
		}
	case 1:
		for i := range lines {
			Sources = append(Sources, models.Protein(lines[i][0]))
		}
	}
}

func FillRoutes(lines [][]string) {
	// Some preprocessing based on the input worked on
	regex := regexp.MustCompile(`^\d{2}-\d{2}-\d{2}$`)

	for i := range lines {
		a := models.Protein(lines[i][0])
		b := models.Protein(lines[i][1])

		if a == "" || b == "" || regex.MatchString(string(a)) || regex.MatchString(string(b)) {
			continue
		}

		AllProteins.Append(a)
		AllProteins.Append(b)
	}

	AllProteins.RemoveDuplicates()
	for _, protein := range AllProteins {
		Routes[protein] = models.ProteinList{}
	}

	for i := range lines {
		a := models.Protein(lines[i][0])
		b := models.Protein(lines[i][1])

		if a == "" || b == "" || regex.MatchString(string(a)) || regex.MatchString(string(b)) {
			continue
		}

		Routes.Insert(a, b)
	}
}

func ReadFile(fileName string, fillVar int, ch chan bool) {
	defer func() {
		ch <- true
	}()
	var csvReader *csv.Reader

	file, err := os.Open(fileName)

	defer file.Close()

	if err != nil {
		fmt.Printf("Couldn't read %q file\n", fileName)
		os.Exit(1)
	}

	csvReader = csv.NewReader(file)
	lines, err := csvReader.ReadAll()

	if err != nil {
		fmt.Printf("Couldn't parse %q\n", fileName)
		os.Exit(1)
	}

	switch fillVar {
	case 0:
		FillRoutes(lines)
	case 1:
		FillVar(lines, fillPlasma)
	case 2:
		FillVar(lines, fillSource)
	}
}
