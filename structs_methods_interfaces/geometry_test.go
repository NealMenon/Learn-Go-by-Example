package main

import "testing"

func TestGeometry(t *testing.T) {
	checker := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("test_Perimeter()", func(t *testing.T) {
		got := Perimeter(11.5, 12.5)
		want := 48.0
		checker(t, got, want)
	})
}
