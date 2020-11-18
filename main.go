package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"./circ"
)

func main() {
	circ.CheckoutStats()
	rows := readSample()
	fmt.Println(rows)
	appendSum(rows)
	writeChanges(rows)
}

func readSample() [][]string {
	f, err := os.Open("data/sample_data.csv")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func appendSum(rows [][]string) {
	rows[0] = append(rows[0], "SUM")
	for i := 1; i < len(rows); i++ {
		rows[i] = append(rows[i], sum(rows[i]))
	}
}

func sum(row []string) string {
	sum := 0
	for _, s := range row {
		x, err := strconv.Atoi(s)
		if err != nil {
			return "NA"
		}
		sum += x
	}
	return strconv.Itoa(sum)
}

func writeChanges(rows [][]string) {
	f, err := os.Create("data/output.csv")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = csv.NewWriter(f).WriteAll(rows)
	if err != nil {
		log.Fatal(err)
	}
}
