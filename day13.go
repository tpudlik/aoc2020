package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/buses"
)

func main() {
	{
		file, err := os.Open("inputs/day13.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		pn, err := buses.ParseSchedules(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Day 13, Part 1: %d\n", buses.Part1(pn))
	}
	{
		file, err := os.Open("inputs/day13.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		c, err := buses.ParseCongruences(file)
		if err != nil {
			log.Fatal(err)
		}
		sol, err := buses.SolveCongruences(c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Day 13, Part 2: %d\n", sol)
	}

}
