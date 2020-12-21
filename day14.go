package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/docking"
)

func main() {
	{
		file, err := os.Open("inputs/day14.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		c := docking.NewComputer()
		if err := docking.ExecuteProgram(file, c); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Day 14, Part 1: %d\n", c.MemorySum())
	}
	{
		file, err := os.Open("inputs/day14.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		c := docking.NewMemoryAddressDecoderChip()
		if err := docking.ExecuteProgram(file, c); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Day 14, Part 2: %d\n", c.MemorySum())
	}
}
