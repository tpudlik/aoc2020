package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/boarding"
)

func main() {
	file, err := os.Open("inputs/day05.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seats []boarding.Seat
	for scanner.Scan() {
		txt := scanner.Text()
		s, err := boarding.DecodePass(txt)
		if err != nil {
			log.Fatal(err)
		}
		seats = append(seats, s)
	}

	highest := 0
	for _, s := range seats {
		if id := s.ID(); id > highest {
			highest = id
		}
	}
	fmt.Printf("Part 1: %d\n", highest)

	seats_map := map[int]bool{}
	for _, s := range seats {
		seats_map[s.ID()] = true
	}
	for i := 1; i < 1<<10; i++ {
		if !seats_map[i] && seats_map[i-1] && seats_map[i+1] {
			fmt.Printf("Part 2: %d\n", i)
			return
		}
	}
}
