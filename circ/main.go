package circ

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var circStats string = "data/2019_circ_stats.csv"

// CirculationStats prints circulation stats
func CirculationStats(year int, month string) {
	fmt.Println(getLocalCircStats(year, month))
}

func getLocalCircStats(year int, month string) []string {
	var circStatsFile string
	var output []string

	circStatsFile = "data/" + strconv.Itoa(year) + "_circ_stats.csv"
	stats := readData(circStatsFile)

	for i := 0; i < len(stats); i++ {
		if stats[i][0] == month {
			output = stats[i]
		}
	}

	if output == nil {
		return []string{"No available data"}
	}
	return output
}

func readData(fileLocation string) [][]string {
	f, err := os.Open(fileLocation)
	defer f.Close()

	if err != nil {
		return [][]string{{"No data available"}}
	}

	rows, _ := csv.NewReader(f).ReadAll()
	return rows
}
