package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/bags"
)

func main() {
	file, err := os.Open("inputs/day07.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules := bags.ParseRules(file)

	fmt.Printf("Part 1: %d\n", rules.ValidContainingBags("shiny gold"))
	fmt.Printf("Part 2: %d\n", rules.NumberOfContainedBags("shiny gold"))
}
