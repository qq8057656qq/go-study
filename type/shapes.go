package _type

import "math"

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Ceil(math.Pi*circle.Radius*circle.Radius*100) / 100
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle Triangle) Area() float64 {
	return triangle.Base * triangle.Height * 0.5
}
