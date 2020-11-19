package main

import (
	"fmt"
	"strings"

	"./circ"
)

func main() {
	fmt.Println("What year do you want data?")
	var year int
	var month string
	fmt.Scanf("%d", &year)
	fmt.Println("Which month?")
	fmt.Scanf("%s", &month)
	circ.CirculationStats(year, strings.Title(month))
}
