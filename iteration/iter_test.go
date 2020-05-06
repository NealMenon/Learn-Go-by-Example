package iteration

import "testing"

func TestIter(t *testing.T) {
	got := Repeat("a", 4)
	want := "aaaa"
	if got != want {
		t.Errorf("Expected %q got %q", want, got)
	}
}
