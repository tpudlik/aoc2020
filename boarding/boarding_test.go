package boarding

import "testing"

func TestPart1Examples(t *testing.T) {
	tests := []struct {
		pass         string
		row, col, id int
	}{
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}
	for _, test := range tests {
		got, err := DecodePass(test.pass)
		if err != nil {
			t.Errorf("Unexpected error: %v for pass %q", err, test.pass)
			continue
		}
		if got.row != test.row {
			t.Errorf("Got row %v, want %v, for pass %q", got.row, test.row, test.pass)
		}
		if got.column != test.col {
			t.Errorf("Got col %v, want %v, for pass %q", got.column, test.col, test.pass)
		}
		if id := got.ID(); id != test.id {
			t.Errorf("Got id %v, want %v, for pass %q", id, test.id, test.pass)
		}
	}
}
