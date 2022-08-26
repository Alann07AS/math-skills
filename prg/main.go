package main

import (
	"fmt"
	"log"
	"math"
	"mathSkills"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		os.Exit(1)
	}
	data, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}
	table := mathSkills.DataToInt(data)
	fmt.Println("Average:", math.Round(mathSkills.Average(table)))
	fmt.Println("Median:", math.Round(mathSkills.Median(table)))
	fmt.Println("Variance:", int(math.Round(mathSkills.Variance(table))))
	fmt.Println("Standard Deviation:", math.Round(mathSkills.Deviation(table)))

}
