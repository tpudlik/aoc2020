package bags

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func ExampleRulesPart1() Rules {
	return ParseRules(strings.NewReader(strings.TrimSpace(`
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`)))
}

func ExampleRulesPart2() Rules {
	return ParseRules(strings.NewReader(strings.TrimSpace(`
shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`)))
}

func TestParseRules(t *testing.T) {
	got := ExampleRulesPart1()
	want := Rules{
		"light red": []Contents{{
			"bright white", 1,
		}, {
			"muted yellow", 2,
		}},
		"dark orange": []Contents{{
			"bright white", 3,
		}, {
			"muted yellow", 4,
		}},
		"bright white": []Contents{{"shiny gold", 1}},
		"muted yellow": []Contents{{"shiny gold", 2}, {"faded blue", 9}},
		"shiny gold":   []Contents{{"dark olive", 1}, {"vibrant plum", 2}},
		"dark olive":   []Contents{{"faded blue", 3}, {"dotted black", 4}},
		"vibrant plum": []Contents{{"faded blue", 5}, {"dotted black", 6}},
		"faded blue":   []Contents{},
		"dotted black": []Contents{},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got\n%v\nWant\n%v\nDiff:\n%v", got, want, cmp.Diff(got, want))
	}
}

func TestValidOutermostBags(t *testing.T) {
	r := ExampleRulesPart1()
	got := r.ValidContainingBags("shiny gold")
	if want := 4; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestNumberOfContainedBags(t *testing.T) {
	tests := []struct {
		rules Rules
		want  int
	}{
		{ExampleRulesPart1(), 32},
		{ExampleRulesPart2(), 126},
	}
	for _, test := range tests {
		got := test.rules.NumberOfContainedBags("shiny gold")
		if got != test.want {
			t.Errorf("got %d, want %d", got, test.want)
		}
	}
}
