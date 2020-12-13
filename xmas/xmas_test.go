package xmas

import (
	"reflect"
	"testing"
)

func TestFirstInvalidNumberShortExamples1(t *testing.T) {
	var list []int
	for i := 1; i < 26; i++ {
		list = append(list, i)
	}

	tests := []struct {
		next int
		want bool
	}{
		{26, true},
		{49, true},
		{100, false},
		{50, false},
	}

	for _, test := range tests {
		testcase := append(list, test.next)
		got, err := FirstInvalidNumber(testcase, 25)
		if test.want != (err != nil) {
			t.Errorf("Unexpected result for case %v: %v, %v", testcase, got, err)
		}
	}
}

func TestFirstInvalidNumberShortExamples2(t *testing.T) {
	var list []int
	for i := 1; i < 26; i++ {
		if i != 20 {
			list = append(list, i)
		}
	}
	list = append(list, 45)

	tests := []struct {
		next int
		want bool
	}{
		{26, true},
		{65, false},
		{64, true},
		{66, true},
	}

	for _, test := range tests {
		testcase := append(list, test.next)
		got, err := FirstInvalidNumber(testcase, 25)
		if test.want != (err != nil) {
			t.Errorf("Unexpected result for case %v: %v, %v", testcase, got, err)
		}
	}
}

func TestFirstInvalidNumberLongExample(t *testing.T) {
	list := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576}
	got, err := FirstInvalidNumber(list, 5)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if want := 127; got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestFirstInvalidNumberFailure(t *testing.T) {
	list := []int{
		204,
		217,
		221,
		234,
		259,
		248,
		250,
		258,
		455,
		275,
		277,
		278,
		296,
		299,
		336,
		330,
		334,
		376,
		346,
		378,
		359,
		379,
		536,
		564,
		533,
		829}
	got, err := FirstInvalidNumber(list, 25)
	if err == nil {
		t.Errorf("Expected no invalid number, got %d", got)
	}
}

func TestContiguousSetSummingTo(t *testing.T) {
	list := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576}
	got, err := ContiguousSetSummingTo(list, 127)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if want := []int{15, 25, 47, 40}; !reflect.DeepEqual(got, want) {
		t.Errorf("Got %d, want %d", got, want)
	}
}
