package structsmethodsandinterfaces

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pow(circle.Radius, 2) * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle Triangle) Area() float64 {
	return triangle.Base * triangle.Height * 0.5
}
