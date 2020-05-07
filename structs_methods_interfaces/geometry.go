package main

// Rectangle is shape with width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns perimeter of rectanlge in float64
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area returns area of quadrilateral float64
func Area(width, height float64) float64 {
	return width * height
}

func main() {

}
