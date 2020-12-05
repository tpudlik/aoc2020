package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/passports"
)

func main() {
	file, err := os.Open("inputs/day04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ps, err := passports.ParseBatch(file)
	if err != nil {
		log.Fatal(err)
	}

	required := 0
	for _, p := range ps {
		if p.RequiredFieldsPresent() {
			required++
		}
	}
	fmt.Printf("Part 1: %v\n", required)

	valid := 0
	for _, p := range ps {
		if p.Valid() {
			valid++
		}
	}
	fmt.Printf("Part 2: %v\n", valid)

}
