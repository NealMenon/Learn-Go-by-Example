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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{10, 15}, 75},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
