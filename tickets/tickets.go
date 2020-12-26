package tickets

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Ticket []int

// A closed interval (one that includes both of its endpoints).
type ClosedInterval struct {
	start, end int
}

type Rule struct {
	name   string
	ranges []ClosedInterval
}

func (r Rule) IsValidForField(field int) bool {
	for _, interval := range r.ranges {
		if field >= interval.start && field <= interval.end {
			return true
		}
	}
	return false
}

type Notes struct {
	rules        []Rule
	myTicket     Ticket
	otherTickets []Ticket
}

func (n Notes) IsTicketValid(t Ticket) bool {
	for _, field := range t {
		valid := false
		for _, r := range n.rules {
			if r.IsValidForField(field) {
				valid = true
				break
			}
		}
		if !valid {
			return false
		}
	}
	return true
}

func (n Notes) TicketScanningErrorRate() int {
	errorRate := 0
	for _, t := range n.otherTickets {
		for _, field := range t {
			valid := false
			for _, r := range n.rules {
				if r.IsValidForField(field) {
					valid = true
					break
				}
			}
			if !valid {
				errorRate += field
			}
		}
	}
	return errorRate
}

func (n Notes) IdentifyFields() (map[string]int, error) {
	candidatesByField := map[string]map[int]bool{}
	for _, rule := range n.rules {
		init := map[int]bool{}
		for i := 0; i < len(n.myTicket); i++ {
			init[i] = true
		}
		candidatesByField[rule.name] = init
	}

	validTickets := []Ticket{}
	for _, t := range n.otherTickets {
		if n.IsTicketValid(t) {
			validTickets = append(validTickets, t)
		}
	}

	for _, t := range validTickets {
		for idx, value := range t {
			for _, r := range n.rules {
				if candidatesByField[r.name][idx] && !r.IsValidForField(value) {
					delete(candidatesByField[r.name], idx)
				}
			}
		}
	}

	// Now we do constraint propagation.
	indexByName := map[string]int{}
	for len(candidatesByField) > 0 {
		var nameToDelete string
		idxToDelete := -1
		for name, candidates := range candidatesByField {
			if len(candidates) == 1 {
				// This field must appear at the given index.  The loop below
				// has only one iteration.
				for idx := range candidates {
					fmt.Printf("%q must appear at index %d\n", name, idx)
					indexByName[name] = idx
					nameToDelete = name
					idxToDelete = idx
				}
				break
			}
		}
		if idxToDelete < 0 {
			return nil, fmt.Errorf("Failed to find a name to delete from map %+v", candidatesByField)
		}
		delete(candidatesByField, nameToDelete)
		for name, candidates := range candidatesByField {
			for idx := range candidates {
				if idx == idxToDelete {
					if len(candidates) == 1 {
						return nil, fmt.Errorf("Problem appears overconstrained: removing index %d from %q after finding it was the only candidate for %q, full map: %+v", idx, name, nameToDelete, candidatesByField)
					}
					delete(candidates, idx)
					candidatesByField[name] = candidates
					break
				}
			}
		}
	}
	return indexByName, nil
}

func (n Notes) ProductOfDepartureFields() (int, error) {
	indexByName, err := n.IdentifyFields()
	if err != nil {
		return 0, err
	}

	out := 1
	for name, idx := range indexByName {
		if !strings.HasPrefix(name, "departure") {
			continue
		}
		out *= n.myTicket[idx]
	}
	return out, nil
}

func (n Notes) MyTicketFieldByIndex(idx int) int {
	return n.myTicket[idx]
}

func ParseNotes(r io.Reader) (Notes, error) {
	scanner := bufio.NewScanner(r)
	notes := Notes{}
	isMyTicket := false
	areOtherTickets := false
	for scanner.Scan() {
		text := scanner.Text()
		if isMyTicket {
			t, err := ParseTicket(text)
			if err != nil {
				return notes, err
			}
			notes.myTicket = t
			isMyTicket = false
			continue
		}
		if areOtherTickets {
			t, err := ParseTicket(text)
			if err != nil {
				return notes, err
			}
			notes.otherTickets = append(notes.otherTickets, t)
			continue
			// We don't set areOtherTickets to false: once we start parsing other tickets, they
			// should constitute the rest of the file.
		}
		if strings.Contains(text, "-") {
			rule, err := ParseRule(text)
			if err != nil {
				return Notes{}, err
			}
			notes.rules = append(notes.rules, rule)
			continue
		}

		if strings.Contains(text, "your ticket:") {
			// The next line contains my ticket
			isMyTicket = true
			continue
		}

		if strings.Contains(text, "nearby tickets:") {
			// The remaining lines contain nearby tickets.
			areOtherTickets = true
			continue
		}
	}
	return notes, scanner.Err()
}

func ParseRule(text string) (Rule, error) {
	pieces := strings.Split(text, ": ")
	if len(pieces) != 2 {
		return Rule{}, fmt.Errorf("Failed to parse rule %q: split on colon produced %d pieces", text, len(pieces))
	}
	name := pieces[0]
	intervals := strings.Split(pieces[1], " or ")
	var parsedIntervals []ClosedInterval
	for _, interval := range intervals {
		bounds := strings.Split(interval, "-")
		if len(bounds) != 2 {
			return Rule{}, fmt.Errorf("Failed to parse %q (in %q) as interval", interval, text)
		}
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			return Rule{}, fmt.Errorf("Failed to parse %q as start of interval: %v", bounds[0], err)
		}
		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			return Rule{}, fmt.Errorf("Failed to parse %q as end of interval: %v", bounds[1], err)
		}
		parsedIntervals = append(parsedIntervals, ClosedInterval{start, end})
	}
	return Rule{name, parsedIntervals}, nil
}

func ParseTicket(text string) (Ticket, error) {
	pieces := strings.Split(text, ",")
	var ticket Ticket
	for _, piece := range pieces {
		i, err := strconv.Atoi(piece)
		if err != nil {
			return ticket, fmt.Errorf("Failed to parse %q as ticket: %v", text, err)

		}
		ticket = append(ticket, i)
	}
	return ticket, nil
}
