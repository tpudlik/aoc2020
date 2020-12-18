package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/seats"
)

func main() {
	{
		file, err := os.Open("inputs/day11.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		s, err := seats.NewFromReader(file, seats.Nearest, 4)
		if err != nil {
			log.Fatal(err)
		}

		s.StepUntilSteadyState()
		fmt.Printf("Day 11 Part 1: %d\n", s.OccupiedSeats())
	}
	{
		file, err := os.Open("inputs/day11.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		s, err := seats.NewFromReader(file, seats.Queen, 5)
		if err != nil {
			log.Fatal(err)
		}

		s.StepUntilSteadyState()
		fmt.Printf("Day 11 Part 2: %d\n", s.OccupiedSeats())
	}
}
