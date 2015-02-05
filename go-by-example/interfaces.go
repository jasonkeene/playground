package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

func (s *square) area() float64 {
	return s.width * s.height
}

func (s *square) perim() float64 {
	return s.width*2 + s.height*2
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return math.Pi * 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func main() {
	s := &square{3, 4}
	c := &circle{5}

	measure(s)
	measure(c)
}
