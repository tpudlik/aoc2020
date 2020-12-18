package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/ship"
)

func main() {
	file, err := os.Open("inputs/day12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions, err := ship.ParseInstructions(file)
	if err != nil {
		log.Fatal(err)
	}

	p := ship.Position{}
	for _, instruction := range instructions {
		p.Update(instruction)
	}
	fmt.Printf("Day 12, Part 1: %d\n", p.ManhattanDistanceTravelled())

	s := ship.NewState()
	for _, instruction := range instructions {
		s.Update(instruction)
	}
	fmt.Printf("Day 12, Part 2: %d\n", s.ManhattanDistanceTravelled())
}
