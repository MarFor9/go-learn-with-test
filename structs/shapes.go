package structs

import "math"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

type Circles struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circles) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}
