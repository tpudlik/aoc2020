package main

import (
	"fmt"
	"log"

	"github.com/tpudlik/aoc2020/password"
)

func main() {
	v, err := password.CountValidPasswordsInFile("inputs/day02.txt", password.Part1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
