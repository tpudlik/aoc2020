package trees

import (
	"strings"
	"testing"
)

func mapForTest() Map {
	r := strings.NewReader(strings.TrimSpace(`
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`))
	return NewMapFromReader(r)
}

func TestCountTrees(t *testing.T) {
	m := mapForTest()
	got := m.CountTreesAlongSlope(1, 3)
	const want int = 7
	if got != want {
		t.Errorf("Encounted %d trees, expected %d", got, want)
	}
}

func TestPart2Examples(t *testing.T) {
	m := mapForTest()
	tests := []struct {
		x, y, want int
	}{
		{1, 1, 2},
		{1, 3, 7},
		{1, 5, 3},
		{1, 7, 4},
		{2, 1, 2},
	}
	for _, test := range tests {
		got := m.CountTreesAlongSlope(test.x, test.y)
		if got != test.want {
			t.Errorf("Got %d, want %d for slope (%d, %d)", got, test.want, test.x, test.y)
		}
	}
}
