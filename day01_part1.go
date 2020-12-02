package main

import (
	"fmt"
	"log"

	"github.com/tpudlik/aoc2020/expense"
)

func main() {
	entries, err := expense.ReadInputFromFile("inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	answer, err := expense.EntriesWithSum(entries, 2020)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)
}
