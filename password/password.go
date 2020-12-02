package password

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Official Toboggan Corporate Policy for passwords.
type Policy struct {
	// Letter to which the policy applies.
	letter rune
	// Minimum number of letter occurrences.
	min int
	// Maximum number of letter occurrences.
	max int
}

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

func (p *Policy) IsSatisfied(password string) bool {
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
}

func ValidateLine(s string) (bool, error) {
	policy_password := strings.Split(s, ": ")
	if len(policy_password) != 2 {
		return false, fmt.Errorf("Can't parse line %q: splitting on ': ' yielded %v", s, policy_password)
	}
	policy, err := NewPolicy(policy_password[0])
	if err != nil {
		return false, err
	}
	return policy.IsSatisfied(policy_password[1]), nil
}

func CountValidPasswords(r io.Reader) (int, error) {
	valid := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		v, err := ValidateLine(txt)
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

func CountValidPasswordsInFile(fname string) (int, error) {
	file, err := os.Open(fname)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return CountValidPasswords(file)
}
