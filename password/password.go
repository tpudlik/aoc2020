package password

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Official Toboggan Corporate Policy for passwords.
type Policy struct {
	// Letter to which the policy applies.
	letter rune
	// Policy parameters with version-dependent interpretation.
	min int
	max int
}

// Indicates whether the policy should be interpreted as in Part 1 or Part 2
// of the problem.
type PolicyVersion int

const (
	// `min` is the minimum number of `letter` occurrences, and `max` is the
	// maximum number of `letter` occurrences.
	Part1 PolicyVersion = iota
	// `letter` must occur at one and only one of index `min` or index `max`.
	// Indexing is 1-based.
	Part2
)

func NewPolicy(s string) (Policy, error) {
	pieces := strings.Split(s, " ")
	if len(pieces) != 2 {
		return Policy{}, fmt.Errorf("Can't decode %q as Policy: can't find letter", s)
	}
	letter := []rune(pieces[1])[0]
	minmax := strings.Split(pieces[0], "-")
	if len(minmax) != 2 {
		return Policy{}, fmt.Errorf("Can't decode %q as Policy: can't decode range", s)
	}
	min, err := strconv.Atoi(minmax[0])
	if err != nil {
		return Policy{}, fmt.Errorf("Can't decode %q as Policy: min is not a number", s)
	}
	max, err := strconv.Atoi(minmax[1])
	if err != nil {
		return Policy{}, fmt.Errorf("Can't decode %q as Policy: max is not a number", s)
	}
	return Policy{letter, min, max}, nil
}

func (p *Policy) IsSatisfied(password string, version PolicyVersion) bool {
	switch version {
	case Part1:
		occurrences := 0
		for _, r := range password {
			if r == p.letter {
				occurrences++
			}
		}
		if occurrences >= p.min && occurrences <= p.max {
			return true
		}
		return false
	case Part2:
		contains_min := false
		contains_max := false
		for i, r := range password {
			if i+1 == p.min && r == p.letter {
				contains_min = true
			}
			if i+1 == p.max && r == p.letter {
				contains_max = true
			}
		}
		return (contains_max && !contains_min) || (!contains_max && contains_min)
	default:
		// Never expected
		log.Fatalf("Invalid version: %d", version)
	}
	return false
}

func ValidateLine(s string, version PolicyVersion) (bool, error) {
	policy_password := strings.Split(s, ": ")
	if len(policy_password) != 2 {
		return false, fmt.Errorf("Can't parse line %q: splitting on ': ' yielded %v", s, policy_password)
	}
	policy, err := NewPolicy(policy_password[0])
	if err != nil {
		return false, err
	}
	return policy.IsSatisfied(policy_password[1], version), nil
}

func CountValidPasswords(r io.Reader, version PolicyVersion) (int, error) {
	valid := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		v, err := ValidateLine(txt, version)
		if err != nil {
			return 0, err
		}
		if v {
			valid++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error scanning input file: %v", err)
	}
	return valid, nil
}

func CountValidPasswordsInFile(fname string, version PolicyVersion) (int, error) {
	file, err := os.Open(fname)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return CountValidPasswords(file, version)
}
