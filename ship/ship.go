package ship

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Heading int

// Possible ship headings.  They're listed in such an order that the default value is East, and
// adding 1 corresponds to rotating right by 90 degrees.
const (
	East Heading = iota
	South
	West
	North
)

type Instruction struct {
	op  rune
	arg int
}

func ParseInstructions(r io.Reader) ([]Instruction, error) {
	instructions := []Instruction{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		op := rune(txt[0])
		arg, err := strconv.Atoi(txt[1:])
		if err != nil {
			return nil, fmt.Errorf("Failed to parse %q as instruction", txt)
		}
		instructions = append(instructions, Instruction{op, arg})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return instructions, nil
}

type Position struct {
	// Latitude and longitude: north and east are positive directions.
	lat, long int
	heading   Heading
}

func (p *Position) Update(instruction Instruction) {
	switch instruction.op {
	case 'N':
		p.lat += instruction.arg
	case 'S':
		p.lat -= instruction.arg
	case 'E':
		p.long += instruction.arg
	case 'W':
		p.long -= instruction.arg
	case 'L':
		p.heading = Heading((int(p.heading) - instruction.arg/90 + 4) % 4)
	case 'R':
		p.heading = Heading((int(p.heading) + instruction.arg/90) % 4)
	case 'F':
		switch p.heading {
		case North:
			p.lat += instruction.arg
		case East:
			p.long += instruction.arg
		case South:
			p.lat -= instruction.arg
		case West:
			p.long -= instruction.arg
		default:
			panic(fmt.Sprintf("Invalid heading: %v", p.heading))
		}
	default:
		panic(fmt.Sprintf("Invalid instruction: %v", instruction))
	}
}

// Computes the absolute value of n.  Oh Go...
//
// There may be more efficient ways:
// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (p *Position) ManhattanDistanceTravelled() int {
	return abs(p.lat) + abs(p.long)
}

type State struct {
	// Ship latitude, longitude.
	s_lat, s_long int

	// Waypoint latitude, longitude, relative to that of the ship.
	w_lat, w_long int
}

func NewState() State {
	return State{0, 0, 1, 10}
}

func (s *State) Update(instruction Instruction) {
	switch instruction.op {
	case 'N':
		s.w_lat += instruction.arg
	case 'S':
		s.w_lat -= instruction.arg
	case 'E':
		s.w_long += instruction.arg
	case 'W':
		s.w_long -= instruction.arg
	case 'L':
		s.RotateWaypointCCW(instruction.arg)
	case 'R':
		s.RotateWaypointCCW(360 - instruction.arg)
	case 'F':
		s.s_lat += s.w_lat * instruction.arg
		s.s_long += s.w_long * instruction.arg
	default:
		panic(fmt.Sprintf("Unexpected instruction: %v", instruction))
	}
}

func (s *State) RotateWaypointCCW(degrees int) {
	switch degrees {
	case 90:
		s.w_lat, s.w_long = s.w_long, -s.w_lat
	case 180:
		s.w_lat, s.w_long = -s.w_lat, -s.w_long
	case 270:
		s.w_lat, s.w_long = -s.w_long, s.w_lat
	default:
		panic(fmt.Sprintf("Unsupported rotation angle: %v degrees", degrees))
	}
}

func (s *State) ManhattanDistanceTravelled() int {
	return abs(s.s_lat) + abs(s.s_long)
}
