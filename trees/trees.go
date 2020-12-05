package trees

import (
	"bufio"
	"io"
)

// Map of an area containing trees.
type Map struct {
	width, height int
	trees         map[Tree]bool
}

type Tree struct {
	x, y int
}

func NewMapFromReader(r io.Reader) Map {
	m := Map{-1, -1, map[Tree]bool{}}
	scanner := bufio.NewScanner(r)
	row := 0
	for scanner.Scan() {
		txt := scanner.Text()
		m.width = len(txt)
		for col, c := range txt {
			if c == '#' {
				m.trees[Tree{row, col}] = true
			}
		}
		row++
	}
	m.height = row
	return m
}

func (m *Map) HasTreeAt(x, y int) bool {
	x_mod := x % m.height
	y_mod := y % m.width
	return m.trees[Tree{x_mod, y_mod}]
}

func (m *Map) CountTreesAlongSlope(vx, vy int) int {
	x, y := 0, 0
	trees := 0
	for x < m.height-1 {
		x += vx
		y += vy
		if m.HasTreeAt(x, y) {
			trees++
		}
	}
	return trees
}
