package password

import (
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	r := strings.NewReader(strings.TrimSpace(`
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`))
	got, err := CountValidPasswords(r, Part1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	want := 2
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestPart2Examples(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", false},
	}

	for _, test := range tests {
		got, err := ValidateLine(test.line, Part2)
		if err != nil {
			t.Errorf("Expected no error, got %v for line %q", err, test.line)
			continue
		}
		if got != test.want {
			t.Errorf("Got %v, want %v for line %q", got, test.want, test.line)
		}
	}
}
