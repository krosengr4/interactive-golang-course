package main

import (
	"math"
)

// Structs for different shapes
type Square struct {
	length float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*(r.length) + 2*(r.width)
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (s Square) Perimeter() float64 {
	return 4 * s.length
}

func (s Square) Area() float64 {
	return math.Pow(s.length, 2)
}

func (c Circle) Perimeter() float64 {
	return c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
