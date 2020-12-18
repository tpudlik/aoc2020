package seats

import (
	"bufio"
	"fmt"
	"io"
)

type Seats struct {
	state      [][]rune
	rows, cols int
	// Which seats are considered neighbors.
	visibility Visibility
	// How many neighboring seats need to be occupied for a seat to transition to empty.
	occupancy_threshold int
}

type position struct {
	idr, idc int
}

// What rule is used to determine which seats are neighbors.
type Visibility int

const (
	// The 8 adjacent squares are neighbors.
	Nearest Visibility = iota
	// Any seat that would be reachable by a chess queen (infinite range in the 8 directions).
	Queen
)

func NewFromReader(r io.Reader, visibility Visibility, occupancy_threshold int) (*Seats, error) {
	s := new(Seats)
	s.visibility = visibility
	s.occupancy_threshold = occupancy_threshold

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		col := []rune{}
		for _, seat := range scanner.Text() {
			col = append(col, seat)
		}
		s.state = append(s.state, col)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(s.state) < 1 {
		return nil, fmt.Errorf("No lines of input successfully read.")
	}
	s.rows = len(s.state)
	s.cols = len(s.state[0])
	return s, nil
}

func (s *Seats) OccupiedSeats() int {
	occupied := 0
	for _, row := range s.state {
		for _, seat := range row {
			if seat == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func (s *Seats) StepUntilSteadyState() int {
	steps := 0
	for s.Step() {
		steps++
	}
	return steps
}

// Advances the state of the seat arrangement by one step.  Returns true if the state of any seat
// has changed.
func (s *Seats) Step() bool {
	next_state := make([][]rune, s.rows)
	for idr := range next_state {
		next_state[idr] = make([]rune, s.cols)
	}

	changed := false
	for idr, row := range s.state {
		for idc, seat := range row {
			on := s.occupiedNeighbors(position{idr, idc})
			if seat == 'L' && on == 0 {
				next_state[idr][idc] = '#'
				changed = true
			} else if seat == '#' && on >= s.occupancy_threshold {
				next_state[idr][idc] = 'L'
				changed = true
			} else {
				next_state[idr][idc] = seat
			}
		}
	}

	s.state = next_state
	return changed
}

func (s *Seats) occupiedNeighbors(pos position) int {
	occupied := 0
	for _, n := range s.neighbors(pos) {
		if n == '#' {
			occupied++
		}
	}
	return occupied
}

func (s *Seats) neighbors(pos position) []rune {
	switch s.visibility {
	case Nearest:
		return s.nearestNeighbors(pos)
	case Queen:
		return s.queenNeighbors(pos)
	default:
		// Not expected.
		panic(fmt.Sprintf("Unsupported visibility: %v", s.visibility))
	}
	return nil
}

func (s *Seats) nearestNeighbors(pos position) []rune {
	row_indices := []int{pos.idr}
	if pos.idr > 0 {
		row_indices = append(row_indices, pos.idr-1)
	}
	if pos.idr < s.rows-1 {
		row_indices = append(row_indices, pos.idr+1)
	}

	col_indices := []int{pos.idc}
	if pos.idc > 0 {
		col_indices = append(col_indices, pos.idc-1)
	}
	if pos.idc < s.cols-1 {
		col_indices = append(col_indices, pos.idc+1)
	}

	// Carthesian product of legal row and col indices.
	neighbors := []rune{}
	for _, idr := range row_indices {
		for _, idc := range col_indices {
			if idr == pos.idr && idc == pos.idc {
				continue
			}
			neighbors = append(neighbors, s.state[idr][idc])
		}
	}
	return neighbors
}

func (s *Seats) queenNeighbors(pos position) []rune {
	// All possible directions on the chessboard.
	directions := []struct {
		x, y int
	}{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
	neighbors := []rune{}
	for _, direction := range directions {
		for steps := 1; ; steps++ {
			idc := pos.idc + direction.x*steps
			idr := pos.idr + direction.y*steps
			if idc >= s.cols || idc < 0 || idr >= s.rows || idr < 0 {
				// Reached the edge of the board without encountering another seat.
				break
			}
			if n := s.state[idr][idc]; n == '#' || n == 'L' {
				// Found the first seat in this direction.
				neighbors = append(neighbors, n)
				break
			}
		}
	}
	return neighbors
}
