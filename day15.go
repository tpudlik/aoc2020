package main

import (
	"fmt"

	"github.com/tpudlik/aoc2020/recitation"
)

func main() {
	fmt.Printf("Day 15, Part 1: %d\n", recitation.Recite([]int{11, 18, 0, 20, 1, 7, 16}, 2020))
	fmt.Printf("Day 15, Part 2: %d\n", recitation.Recite([]int{11, 18, 0, 20, 1, 7, 16}, 30000000))
}
