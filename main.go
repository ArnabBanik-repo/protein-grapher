package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/ArnabBanik-repo/protein-grapher/db"
	"github.com/ArnabBanik-repo/protein-grapher/models"
	"github.com/ArnabBanik-repo/protein-grapher/utils"
)

const (
	saveDirectory     string = "./output"
	mappingDirectory  string = "./output/mappedPlasma"
	unmappedDirectory string = "./output/unmappedPlasma"
	routeFile         int    = 0
	plasmaFile        int    = 1
	sourceFile        int    = 2
)

func main() {
	createSaveDirectory(saveDirectory)
	createSaveDirectory(mappingDirectory)
	createSaveDirectory(unmappedDirectory)

	routeFileName := flag.String("routes", "./db/routes.csv", "Csv file with the routes specified in cols A and B")
	plasmaFileName := flag.String("plasma", "./db/plasma.csv", "Csv file with the plasma proteins specified in col A")
	sourcesFileName := flag.String("sources", "./db/sources.csv", "Csv file with the drug proteins specified in col A")
	flag.Parse()

	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	fmt.Println("Please wait while we are processing ...")

	go db.ReadFile(*routeFileName, routeFile, ch1)
	go db.ReadFile(*plasmaFileName, plasmaFile, ch2)
	go db.ReadFile(*sourcesFileName, sourceFile, ch3)

	<-ch1
	<-ch2
	<-ch3

	for _, i := range db.Plasma {
		db.AllProteins.Append(i)
	}

	var wg sync.WaitGroup
	for _, source := range db.Sources {
		wg.Add(1)
		go func(source models.Protein) {
			utils.ProcessDrug(source)
			wg.Done()
		}(source)
	}

	wg.Wait()

	fmt.Println("Processing Complete!\nPlease check the output folder for the results")
}

func createSaveDirectory(path string) {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("%q directory already exists. Existing files will be overwritten.\nContinue? (Y/n)\n", path)

		var ch string
		fmt.Scanf("%s", &ch)

		if len(ch) > 0 {
			switch string(ch[0]) {
			case "n":
				fmt.Println("Quitting ...")
				os.Exit(1)
			case "N":
				fmt.Println("Quitting ...")
				os.Exit(1)
			}
		}

		fmt.Println()

	} else if !os.IsNotExist(err) {
		fmt.Printf("Error accessing output directory: %v\n", err)
		return
	}

	os.Mkdir(path, os.ModePerm)
}
