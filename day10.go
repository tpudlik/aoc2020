package main

import (
	"fmt"
	"log"

	"github.com/tpudlik/aoc2020/adapters"
	"github.com/tpudlik/aoc2020/input"
)

func main() {
	list, err := input.ReadIntList("inputs/day10.txt")
	if err != nil {
		log.Fatal(err)
	}

	diffs := adapters.GetDiffs(list)
	fmt.Printf("Part 1: %d\n", diffs[1]*diffs[3])
	fmt.Printf("Part 2: %d\n", adapters.CountArragements(list))
}
