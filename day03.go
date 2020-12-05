package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/trees"
)

func main() {
	file, err := os.Open("inputs/day03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := trees.NewMapFromReader(file)
	fmt.Printf("Part 1: %d\n", m.CountTreesAlongSlope(1, 3))

	slopes := []struct{ x, y int }{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	product := 1
	for _, slope := range slopes {
		product *= m.CountTreesAlongSlope(slope.x, slope.y)
	}
	fmt.Printf("Part 2: %d\n", product)
}
