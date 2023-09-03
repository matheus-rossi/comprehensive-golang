package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
}

type square struct {
	side float32
}

type circle struct {
	radius float32
}

func (s square) area() float32 {
	return s.side * s.side
}

func (c circle) area() float32 {
	return c.radius * math.Pi
}

func calculate(g geometry) {
	fmt.Println(g.area())
}

func main() {
	squareOne := square{side: 10}
	circleOne := circle{radius: 10}

	calculate(squareOne)
	calculate(circleOne)

}
