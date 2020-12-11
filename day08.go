package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/handheld"
)

func main() {
	file, err := os.Open("inputs/day08.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	program, err := handheld.ParseProgram(file)
	if err != nil {
		log.Fatal(err)
	}

	acc, err := handheld.DetectInfiniteLoop(program)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1: %d\n", acc)

	acc, err = handheld.FixBootLoop(program)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 2: %d\n", acc)
}
