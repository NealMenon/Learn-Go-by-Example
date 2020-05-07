package main

import "math"

// Rectangle is shape with width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle with radius
type Circle struct {
	Radius float64
}

// Triangle base and height
type Triangle struct {
	base, height float64
}

// Shape interface
type Shape interface {
	Area() float64
}

// Perimeter returns perimeter of rectanlge in float64
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

// Area returns area of quadrilateral float64
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Area of circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Area of triangle
func (t Triangle) Area() float64 {
	return t.height * t.base / 2.0
}

func main() {

}
