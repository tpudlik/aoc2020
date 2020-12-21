package docking

import (
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	program := strings.NewReader(`
		mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
		mem[8] = 11
		mem[7] = 101
		mem[8] = 0`)
	c := NewComputer()
	if err := ExecuteProgram(program, c); err != nil {
		t.Fatal(err)
	}
	got := c.MemorySum()
	if want := int64(165); got != want {
		t.Errorf("Got %d, want %d for memory sum.  Decoder chip state: %+v", got, want, c)
	}
}

func TestPart2Example(t *testing.T) {
	program := strings.NewReader(`
		mask = 000000000000000000000000000000X1001X
		mem[42] = 100
		mask = 00000000000000000000000000000000X0XX
		mem[26] = 1`)
	c := NewMemoryAddressDecoderChip()
	if err := ExecuteProgram(program, c); err != nil {
		t.Fatal(err)
	}
	got := c.MemorySum()
	if want := int64(208); got != want {
		t.Errorf("Got %d, want %d for memory sum.  Decoder chip state: %+v", got, want, c)
	}
}
