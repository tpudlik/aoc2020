package ship

import (
	"strings"
	"testing"
)

func TestPositionUpdate(t *testing.T) {
	tests := []struct {
		instruction Instruction
		pos         Position
	}{
		{Instruction{'F', 10}, Position{0, 10, East}},
		{Instruction{'N', 3}, Position{3, 10, East}},
		{Instruction{'F', 7}, Position{3, 17, East}},
		{Instruction{'R', 90}, Position{3, 17, South}},
		{Instruction{'F', 11}, Position{-8, 17, South}},
		{Instruction{'L', 90}, Position{-8, 17, East}},
		{Instruction{'L', 90}, Position{-8, 17, North}},
		{Instruction{'L', 90}, Position{-8, 17, West}},
		{Instruction{'L', 90}, Position{-8, 17, South}},
		{Instruction{'R', 90}, Position{-8, 17, West}},
		{Instruction{'R', 90}, Position{-8, 17, North}},
		{Instruction{'R', 90}, Position{-8, 17, East}},
		{Instruction{'R', 90}, Position{-8, 17, South}},
		{Instruction{'R', 90}, Position{-8, 17, West}},
	}
	p := Position{}
	for _, test := range tests {
		p.Update(test.instruction)
		if p.long != test.pos.long || p.lat != test.pos.lat || p.heading != test.pos.heading {
			t.Errorf("Got position %v, want %v after instruction %v", p, test.pos, test.instruction)
			break
		}
	}
}

func TestPart1Example(t *testing.T) {
	instructions, err := ParseInstructions(strings.NewReader(strings.TrimSpace(`
F10
N3
F7
R90
F11
`)))
	if err != nil {
		t.Fatal(err)
	}

	pos := Position{}
	for _, instruction := range instructions {
		pos.Update(instruction)
	}
	got := pos.ManhattanDistanceTravelled()
	if want := 25; got != want {
		t.Errorf("Got %d, want %d for distance travelled; final position %v", got, want, pos)
	}
}

func TestPart2Example(t *testing.T) {
	instructions, err := ParseInstructions(strings.NewReader(strings.TrimSpace(`
F10
N3
F7
R90
F11
`)))
	if err != nil {
		t.Fatal(err)
	}

	s := NewState()
	for _, instruction := range instructions {
		s.Update(instruction)
	}
	got := s.ManhattanDistanceTravelled()
	if want := 286; got != want {
		t.Errorf("Got %d, want %d for distance travelled; final state %v", got, want, s)
	}
}
