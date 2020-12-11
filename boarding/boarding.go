package boarding

import "fmt"

type Seat struct {
	row, column int
}

func (s Seat) ID() int {
	return s.row*8 + s.column
}

func DecodePass(pass string) (Seat, error) {
	if l := len(pass); l != 10 {
		return Seat{}, fmt.Errorf("Wrong pass length %d, expected 10", l)
	}
	row := 0
	for i := 0; i < 7; i++ {
		if pass[i] == 'B' {
			row += 1 << (6 - i)
		}
	}
	column := 0
	for i := 7; i < 10; i++ {
		if pass[i] == 'R' {
			column += 1 << (9 - i)
		}
	}
	return Seat{row, column}, nil
}
