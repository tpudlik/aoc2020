package expense

import (
	"reflect"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	entries := []int{1721,
		979,
		366,
		299,
		675,
		1456}
	got, err := EntriesWithSum(entries, 2020)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	want := 514579
	if got != want {
		t.Errorf("Got %d, want %d for entries %v", got, want, entries)
	}
}

func TestReadInput(t *testing.T) {
	r := strings.NewReader(strings.TrimSpace(`
123
456
789`))
	got, err := ReadInput(r)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	want := []int{123, 456, 789}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
