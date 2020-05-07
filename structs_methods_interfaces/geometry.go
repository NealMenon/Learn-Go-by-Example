package main

// Rectangle is shape with width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns perimeter of rectanlge in float64
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

// Area returns area of quadrilateral float64
func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}

func main() {

}
