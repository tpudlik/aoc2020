package tickets

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	n, err := ParseNotes(strings.NewReader(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`))
	if err != nil {
		t.Fatal(err)
	}
	got := n.TicketScanningErrorRate()
	if want := 71; got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestPart2Example(t *testing.T) {
	n, err := ParseNotes(strings.NewReader(`class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`))
	if err != nil {
		t.Fatal(err)
	}
	got, err := n.IdentifyFields()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]int{"row": 0, "class": 1, "seat": 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %+v, want %+v", got, want)
	}
}
