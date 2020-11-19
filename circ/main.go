package circ

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var circStats string = "data/2019_circ_stats.csv"

// CirculationStats prints circulation stats
func CirculationStats(year int, month string) {
	fmt.Println(getLocalCircStats(year, month))
}

func getLocalCircStats(year int, month string) []string {
	var circStatsFile string
	var output []string
	switch year {
	case 2019:
		circStatsFile = "data/2019_circ_stats.csv"
	default:
		return []string{"No data for the requested year"}
	}
	stats := readData(circStatsFile)

	for i := 0; i < len(stats); i++ {
		if stats[i][0] == month {
			output = stats[i]
		}
	}

	return output
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
