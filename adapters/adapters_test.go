package adapters

import "testing"

func TestSmallExample(t *testing.T) {
	list := []int{16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4}
	got := CountArragements(list)
	want := 8
	if got != want {
		t.Errorf("got %d != %d", got, want)
	}
}

func TestLongExample(t *testing.T) {
	list := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3}
	got := CountArragements(list)
	want := 19208
	if got != want {
		t.Errorf("got %d != %d", got, want)
	}
}
