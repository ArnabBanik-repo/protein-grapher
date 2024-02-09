package utils

import (
	"fmt"

	"github.com/ArnabBanik-repo/protein-grapher/db"
	"github.com/ArnabBanik-repo/protein-grapher/models"
)

func ProcessDrug(drug models.Protein) {
  x := db.Routes.ShortestPaths(drug)
  paths := []models.ProteinList{}
	mappedPlasma := models.ProteinList{}
	unmappedPlasma := []models.ProteinList{}

	for _, r := range x {
		if len(r) > 0 && db.Plasma.ContainsProtein(r[len(r)-1]) {
			paths = append(paths, r)
			mappedPlasma = append(mappedPlasma, r[len(r)-1])
		}
	}

	for _, p := range db.Plasma {
		if !mappedPlasma.ContainsProtein(p) {
			unmappedPlasma = append(unmappedPlasma, models.ProteinList{p})
		}
	}

	if len(paths) == 0 {
		fmt.Println(drug, "❌")
		return
	}

	// saveToExcel(paths, drug)
	saveToCsv(paths, drug, "mappedPlasma")
	saveToCsv(unmappedPlasma, drug, "unmappedPlasma")
}
