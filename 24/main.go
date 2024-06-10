package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.
func main() {

	// точки для примера
	A := NewPoint(1, 1)
	B := NewPoint(3, 4)

	// расстояние между ними
	distance := A.Distance(*B)

	fmt.Println(distance)
}

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt(float64((p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y)))
}

func (p Point) GetX() int {
	return p.x
}

func (p Point) GetY() int {
	return p.y
}
