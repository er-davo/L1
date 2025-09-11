package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(
		q.x-p.x,
		q.y-p.y,
	)
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func main() {
	var x, y float64

	fmt.Printf("Enter x and y for first point: ")
	fmt.Scan(&x, &y)

	p1 := NewPoint(x, y)

	fmt.Printf("Enter x and y for second point: ")
	fmt.Scan(&x, &y)

	p2 := NewPoint(x, y)

	fmt.Println("Distance:", p1.Distance(p2))
}
