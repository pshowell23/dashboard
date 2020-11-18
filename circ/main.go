package circ

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// CheckoutStats prints message
func CheckoutStats() {
	circStats := "data/2019_circ_stats.csv"
	fmt.Println("Checkout stats")
	fmt.Println(readData(circStats))
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
