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
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rectangle_Area", shape: Rectangle{Width: 12, Height: 6}, area: 72.0},
		{name: "Circle_Area", shape: Circle{Radius: 10}, area: 314.1592653589793},
		{name: "Triangle_Area", shape: Triangle{Base: 12, Height: 6}, area: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.area {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.area)
			}
		})

	}
}
