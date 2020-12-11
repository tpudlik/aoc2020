package bags

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Contents struct {
	Color  string
	Number int
}

// The bag content rules, represented as a mapping from bag color to the slice
// of (color, number) of allowed contained bags.  Basically a weighted directed
// graph.
type Rules map[string][]Contents

// Returns the number of bag colors that may be contained in outermost.
func (r Rules) ValidContainedBags(outermost string) int {
	valid := map[string]bool{}
	frontier := r[outermost]
	var current Contents
	for len(frontier) > 0 {
		current, frontier = frontier[0], frontier[1:]
		valid[current.Color] = true
		for _, neighbor := range r[current.Color] {
			if _, ok := valid[neighbor.Color]; ok {
				// We've already considered this bag color.
				continue
			}
			frontier = append(frontier, neighbor)
		}
	}
	return len(valid)
}

// Returns the number of bags that may contain innermost.
func (r Rules) ValidContainingBags(innermost string) int {
	// Map from bag color to color of allowed directly containing bags.
	inverse := map[string][]string{}

	for color, contents := range r {
		for _, contained := range contents {
			inverse[contained.Color] = append(inverse[contained.Color], color)
		}
	}

	valid := map[string]bool{}
	frontier := inverse[innermost]
	var current string
	for len(frontier) > 0 {
		current, frontier = frontier[0], frontier[1:]
		valid[current] = true
		for _, neighbor := range inverse[current] {
			if _, ok := valid[neighbor]; ok {
				// We've already considered this bag color.
				continue
			}
			frontier = append(frontier, neighbor)
		}
	}
	return len(valid)
}

// Returns the number of bags that outermost will contain.
func (r Rules) NumberOfContainedBags(outermost string) int {
	total := 0
	frontier := r[outermost]
	var current Contents
	for len(frontier) > 0 {
		current, frontier = frontier[0], frontier[1:]
		total += current.Number
		for _, neighbor := range r[current.Color] {
			required := Contents{
				Color:  neighbor.Color,
				Number: neighbor.Number * current.Number,
			}
			frontier = append(frontier, required)
		}
	}
	return total
}

var ruleRe = regexp.MustCompile(`(?P<color>[a-z A-Z]+) bags contain (?P<contents>[0-9a-zA-Z ,]+)`)
var contentRe = regexp.MustCompile(`(?P<number>[0-9]+) (?P<color>[a-zA-Z ]+) bag`)

func ParseRules(r io.Reader) Rules {
	rules := Rules{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()

		matches := ruleRe.FindStringSubmatch(txt)
		if matches == nil {
			panic(fmt.Sprintf("%q did not match ruleRe %v", txt, ruleRe))
		}
		color := matches[ruleRe.SubexpIndex("color")]
		contents := matches[ruleRe.SubexpIndex("contents")]

		contained_bags := []Contents{}
		for _, content := range strings.Split(contents, ", ") {
			if content == "no other bags" {
				// Special case
				continue
			}
			m := contentRe.FindStringSubmatch(content)
			if m == nil {
				// Not expected
				panic(fmt.Sprintf("%q does not match contentRe", content))
			}
			number, err := strconv.Atoi(m[contentRe.SubexpIndex("number")])
			if err != nil {
				panic(err)
			}
			c := Contents{
				m[contentRe.SubexpIndex("color")],
				number,
			}
			contained_bags = append(contained_bags, c)
		}
		rules[color] = contained_bags
	}
	return rules
}
