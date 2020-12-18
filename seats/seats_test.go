package seats

import (
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	s, err := NewFromReader(strings.NewReader(strings.TrimSpace(`
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`)), Nearest, 4)
	if err != nil {
		t.Fatalf("Unexpected error constructing Seats instance: %v", err)
	}
	steps := s.StepUntilSteadyState()
	if want := 5; steps != want {
		t.Errorf("Took %d steps to reach steady state, want %d", steps, want)
	}
	occupied := s.OccupiedSeats()
	if want := 37; occupied != want {
		t.Errorf("Found %d occupied seats, want %d", occupied, want)
	}
}

func TestPart2Example(t *testing.T) {
	s, err := NewFromReader(strings.NewReader(strings.TrimSpace(`
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`)), Queen, 5)
	if err != nil {
		t.Fatalf("Unexpected error constructing Seats instance: %v", err)
	}
	steps := s.StepUntilSteadyState()
	if want := 6; steps != want {
		t.Errorf("Took %d steps to reach steady state, want %d", steps, want)
	}
	occupied := s.OccupiedSeats()
	if want := 26; occupied != want {
		t.Errorf("Found %d occupied seats, want %d", occupied, want)
	}
}
