package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpudlik/aoc2020/tickets"
)

func main() {
	file, err := os.Open("inputs/day16.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	n, err := tickets.ParseNotes(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 16, Part 1: %d\n", n.TicketScanningErrorRate())

	product, err := n.ProductOfDepartureFields()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 16, Part 2: %d\n", product)
}
