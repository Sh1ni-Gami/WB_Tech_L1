package main

import (
	"fmt"
	"math"
)

// структура для представления точки
type Point struct {
	x float64
	y float64
}

// конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// вычисляет расстояние между двумя точками
func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	point1 := NewPoint(3, 4)
	point2 := NewPoint(7, 1)

	distance := point1.Distance(point2)

	fmt.Printf("Расстояние между точками (%.2f, %.2f) и (%.2f, %.2f) составляет: %.2f\n",
		point1.x, point1.y, point2.x, point2.y, distance)
}
