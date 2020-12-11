package handheld

import (
	"strings"
	"testing"
)

func TestInfiniteLoopExample(t *testing.T) {
	program, err := ParseProgram(strings.NewReader(strings.TrimSpace(`
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`)))
	if err != nil {
		t.Fatalf("Program parse error: %q", err)
	}

	got, err := DetectInfiniteLoop(program)
	if err != nil {
		t.Fatalf("Unexpected error: %q", err)
	}
	if got != 5 {
		t.Errorf("got %d, want %d", got, 5)
	}
}

func TestFixBootLoop(t *testing.T) {
	program, err := ParseProgram(strings.NewReader(strings.TrimSpace(`
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`)))
	if err != nil {
		t.Fatalf("Program parse error: %q", err)
	}

	got, err := FixBootLoop(program)
	if err != nil {
		t.Fatalf("Unexpected error: %q", err)
	}
	if got != 8 {
		t.Errorf("got %d, want %d", got, 8)
	}
}
