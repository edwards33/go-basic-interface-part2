package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape){
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	newSquare := square{side: 20}
	newCircle := circle{radius: 15.5}
	info(newSquare)
	info(newCircle)
}
