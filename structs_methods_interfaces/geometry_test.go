package main

import "testing"

func TestGeometry(t *testing.T) {
	recty := Rectangle{11.5, 12.5}
	checker := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("test_Perimeter()", func(t *testing.T) {
		got := Perimeter(recty)
		want := 48.0
		checker(t, got, want)
	})
	t.Run("test_Area()", func(t *testing.T) {
		got := Area(recty)
		want := 143.75
		checker(t, got, want)
	})
}
