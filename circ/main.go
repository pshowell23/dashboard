package circ

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var circStats string = "data/2019_circ_stats.csv"

// CheckoutStats prints circulation stats
func CheckoutStats() {
	fmt.Println(getCircStats())
}

func getCircStats() [][]string {
	return readData(circStats)
}

func readData(fileLocation string) [][]string {
	f, err := os.Open(fileLocation)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, _ := csv.NewReader(f).ReadAll()
	return rows
}
