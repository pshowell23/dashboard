package circ

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// CirculationStats prints circulation stats
func CirculationStats(year int, month string) {
	localCircStats := getLocalCircStats(year, month)
	illBorrowStats := getILLBorrowStats(year, month)
	if len(localCircStats) == 5 {
		fmt.Println("For the month of", localCircStats[0], year)
		fmt.Println("Items Checked Out:", localCircStats[1])
		fmt.Println("Items Checked In:", localCircStats[2])
		fmt.Println("Items Renewed:", localCircStats[3])
		fmt.Println("Items Soft Checked Out:", localCircStats[4])
	} else {
		fmt.Println(localCircStats[0])
	}
	if len(illBorrowStats) == 3 {
		fmt.Println("Borrowed Articles:", illBorrowStats[1][1])
		fmt.Println("Borrowed Books:", illBorrowStats[2][1])
	} else {
		fmt.Println(illBorrowStats[0][0])
	}
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

func getILLBorrowStats(year int, month string) [][]string {
	var illBorrowFile string

	illBorrowFile = "data/" + strconv.Itoa(year) + "_" + month + "_ILL_borrow_requests.csv"

	borrowStats := readData(illBorrowFile)

	if borrowStats == nil {
		log.Fatal("No Data")
	}
	return borrowStats
}

func readData(fileLocation string) [][]string {
	f, err := os.Open(fileLocation)
	defer f.Close()

	if err != nil {
		return [][]string{{"No ILL data available"}}
	}

	rows, _ := csv.NewReader(f).ReadAll()
	return rows
}
