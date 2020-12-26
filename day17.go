package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/cubes"
)

func solveForDimension(d cubes.Dimensions) int {
	file, err := os.Open("inputs/day17.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	g, err := cubes.ParseGrid(file)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 6; i++ {
		g.Step(d)
	}
	return g.CountActive()
}

func main() {
	fmt.Printf("Day 17, Part 1: %d\n", solveForDimension(cubes.Three))
	fmt.Printf("Day 17, Part 2: %d\n", solveForDimension(cubes.Four))
}
