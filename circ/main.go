package circ

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// CirculationStats prints circulation stats
func CirculationStats(year int, month string) {
	stats := getLocalCircStats(year, month)
	fmt.Println("For the month of", stats[0], year)
	fmt.Println("Items Checked Out:", stats[1])
	fmt.Println("Items Checked In:", stats[2])
	fmt.Println("Items Renewed:", stats[3])
	fmt.Println("Items Soft Checked Out:", stats[4])
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
