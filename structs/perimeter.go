package structs

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float32
}

type Triangle struct {
	Base float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return (r.Height + r.Width) * 2
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}
