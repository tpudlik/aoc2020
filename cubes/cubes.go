package cubes

import (
	"bufio"
	"fmt"
	"io"
)

type Dimensions int

const (
	Three Dimensions = 3
	Four             = 4
)

type Coordinates struct {
	x, y, z, w int
}

type Grid struct {
	active map[Coordinates]bool
}

func ParseGrid(r io.Reader) (*Grid, error) {
	scanner := bufio.NewScanner(r)
	y := 0
	g := Grid{active: map[Coordinates]bool{}}
	for scanner.Scan() {
		for x, state := range scanner.Text() {
			if state == '#' {
				g.active[Coordinates{x, y, 0, 0}] = true
			}
		}
		y++
	}
	return &g, scanner.Err()
}

func (g *Grid) Step(d Dimensions) {
	newActive := map[Coordinates]bool{}
	activeByCoords := g.ActiveAndNeighbors(d)
	for coords, active := range activeByCoords {
		neighbors := Neighbors(coords, d)
		activeNeighbors := 0
		for _, n := range neighbors {
			if activeByCoords[n] {
				activeNeighbors++
			}
		}
		if active && ((activeNeighbors == 2) || (activeNeighbors == 3)) {
			newActive[coords] = true
		}
		if !active && (activeNeighbors == 3) {
			newActive[coords] = true
		}
	}
	g.active = newActive
}

// Returns a map from grid coordinates to state (true for active).  The map includes all active
// grid coordinates, and all of their neighbors.
func (g *Grid) ActiveAndNeighbors(d Dimensions) map[Coordinates]bool {
	out := map[Coordinates]bool{}
	for coords, active := range g.active {
		out[coords] = active
		for _, n := range Neighbors(coords, d) {
			out[n] = g.active[n]
		}
	}
	return out
}

func (g *Grid) CountActive() int {
	out := 0
	for _, active := range g.active {
		if active {
			out++
		}
	}
	return out
}

func Neighbors(q Coordinates, d Dimensions) []Coordinates {
	var dwStart, dwEnd int
	switch d {
	case Three:
		dwStart = 0
		dwEnd = 0
	case Four:
		dwStart = -1
		dwEnd = 1
	default:
		panic(fmt.Sprintf("Unsupported dimension number: %v", d))
	}
	out := []Coordinates{}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := dwStart; dw <= dwEnd; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					out = append(out, Coordinates{q.x + dx, q.y + dy, q.z + dz, q.w + dw})
				}
			}
		}
	}
	return out
}
