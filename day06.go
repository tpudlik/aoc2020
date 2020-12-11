package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Questions to which a group answered yes.
type Group map[rune]bool

// If "anyone" is true, then a question is considered to be answered yes if
// _anyone_ in the group answered yes.  Otherwise, it's considered to be
// answered yes only if _everyone_ in the group answered yes.
func ParseGroups(anyone bool) []Group {
	file, err := os.Open("inputs/day06.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	groups := []Group{}

	in_group := false
	current := Group{}
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.TrimSpace(txt) == "" {
			if in_group {
				// End of a group.
				groups = append(groups, current)
				in_group = false
			}
			// Nothing more to do.
			continue
		}
		new_group := false
		if !in_group {
			// Start a new group
			current = Group{}
			in_group = true
			new_group = true
		}

		if anyone || new_group {
			for _, c := range txt {
				current[c] = true
			}
		} else {
			del := []rune{}
			for k := range current {
				if !strings.ContainsRune(txt, k) {
					del = append(del, k)
				}
			}
			for _, k := range del {
				delete(current, k)
			}
		}
	}
	// If input doesn't end with an empty line, we need to close the last
	// group.
	if in_group {
		groups = append(groups, current)
	}

	return groups
}

func main() {
	// Part 1
	groups := ParseGroups(true)
	counts := 0
	for _, g := range groups {
		counts += len(g)
	}
	fmt.Printf("Part 1: %d\n", counts)

	// Part 2
	groups = ParseGroups(false)
	counts = 0
	for _, g := range groups {
		counts += len(g)
	}
	fmt.Printf("Part 2: %d\n", counts)
}
