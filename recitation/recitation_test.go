package recitation

import "testing"

func TestExtendedExample(t *testing.T) {
	want_spoken := []int{0, 3, 6, 0, 3, 3, 1, 0, 4, 0}
	for idx, want := range want_spoken {
		got := Recite([]int{0, 3, 6}, idx+1)
		if got != want {
			t.Errorf("Spoken number %d: got %d, want %d", idx+1, got, want)
		}
	}
}

func Test2020thNumberInPart1Example(t *testing.T) {
	got := Recite([]int{0, 3, 6}, 2020)
	if want := 436; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFewMorePart1Examples(t *testing.T) {
	tests := []struct {
		starting []int
		want     int
	}{
		{[]int{1, 3, 2}, 1},
		{[]int{2, 1, 3}, 10},
		{[]int{1, 2, 3}, 27},
		{[]int{2, 3, 1}, 78},
		{[]int{3, 2, 1}, 438},
		{[]int{3, 1, 2}, 1836},
	}

	for _, test := range tests {
		got := Recite(test.starting, 2020)
		if got != test.want {
			t.Errorf("Starting numbers %v: got %d, want %d", test.starting, got, test.want)
		}
	}
}

func BenchmarkReciteRef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReciteRef([]int{3, 1, 2}, 20000)
	}
}

func BenchmarkRecite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recite([]int{3, 1, 2}, 20000)
	}
}
