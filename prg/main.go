package main

import (
	"fmt"
	"log"
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
	fmt.Println(table)
	fmt.Println("Average:", mathSkills.Average(table))
	fmt.Println("Median:", mathSkills.Median(table))
	fmt.Println("Variance:", mathSkills.Variance(table))
	fmt.Println("Standard Deviation:", mathSkills.Deviation(table))

}
