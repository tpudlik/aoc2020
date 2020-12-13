package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tpudlik/aoc2020/xmas"
)

func min(vars []int) int {
	min := vars[0]

	for _, i := range vars {
		if i < min {
			min = i
		}
	}
	return min
}

func max(vars []int) int {
	max := vars[0]

	for _, i := range vars {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {
	file, err := os.Open("inputs/day09.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		list = append(list, n)
	}

	n, err := xmas.FirstInvalidNumber(list, 25)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", n)

	s, err := xmas.ContiguousSetSummingTo(list, n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d\n", max(s)+min(s))
}
