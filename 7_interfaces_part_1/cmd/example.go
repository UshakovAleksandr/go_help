package cmd

import "math"

type Shape interface {
	Area() float64
}

type Square struct {
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func NewSquare(side float64) Shape {
	return &Square{side: side}
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func NewCircle(radius float64) Shape {
	return &Circle{radius: radius}
}

func ShapeArea(shape Shape) float64 {
	return shape.Area()
}

