package cubes

import (
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		d    Dimensions
		want int
	}{
		{Three, 112},
		{Four, 848},
	}
	for _, test := range tests {
		g, err := ParseGrid(strings.NewReader(`.#.
..#
###`))
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 6; i++ {
			g.Step(test.d)
		}
		got := g.CountActive()
		if got != test.want {
			t.Errorf("Case %v: Got %d, want %d; state: %+v", test.d, got, test.want, g.active)
		}
	}
}
