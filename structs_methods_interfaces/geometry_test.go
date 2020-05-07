package main

import "testing"

func TestPerimeter(t *testing.T) {
	checker := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("rectangles_perimeter", func(t *testing.T) {
		recty := Rectangle{11.5, 12.5}
		got := Perimeter(recty)
		want := 48.0
		checker(t, got, want)
	})
}

func TestArea(t *testing.T) {
	checker := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("rectangles_area", func(t *testing.T) {
		recty := Rectangle{11.5, 12.5}
		got := recty.Area()
		want := 143.75
		checker(t, got, want)
	})
	t.Run("circles", func(t *testing.T) {
		circy := Circle{10}
		got := circy.Area()
		want := 314.1592653589793
		checker(t, got, want)
	})
}
