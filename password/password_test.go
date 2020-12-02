package password

import (
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	r := strings.NewReader(strings.TrimSpace(`
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`))
	got, err := CountValidPasswords(r)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	want := 2
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
