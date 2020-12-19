package buses

import (
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	r := strings.NewReader(strings.TrimSpace(`
939
7,13,x,x,59,x,31,19
`))
	pn, err := ParseSchedules(r)
	if err != nil {
		t.Fatal(err)
	}
	got := Part1(pn)
	if want := 295; got != want {
		t.Errorf("Got %d, want %d for notes %v", got, want, pn)
	}
}

func TestPart2Examples(t *testing.T) {
	tests := []struct {
		s    string
		want int64
	}{
		{`
939
7,13,x,x,59,x,31,19
		`,
			1068781},
		{`
939
17,x,13,19
`,
			3417,
		}, {`
939
67,7,59,61
		`,
			754018,
		}, {`
939
67,x,7,59,61
`,
			779210,
		}, {`
939
67,7,x,59,61
`,
			1261476,
		}, {`
939
1789,37,47,1889
`,
			1202161486,
		},
	}
	for idx, test := range tests {
		r := strings.NewReader(strings.TrimSpace(test.s))
		c, err := ParseCongruences(r)
		if err != nil {
			t.Errorf("Case %d: Failed to parse congruences: %v", idx, err)
			continue
		}
		got, err := SolveCongruences(c)
		if err != nil {
			t.Errorf("Case %d: Failed to solve congruences: %v", idx, err)
			continue
		}
		if got != test.want {
			t.Errorf("Case %d: Got %d, want %d", idx, got, test.want)
		}
	}
}
