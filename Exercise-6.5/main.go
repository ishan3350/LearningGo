package main

import (
	"fmt"
	"math"
)

type square struct {
	a int
}

type circle struct {
	r int
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return float64(s.a * s.a)
}

func (c circle) area() float64 {
	r := float64(c.r)
	return math.Pi * (r*r)
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	s1 := square{
		a: 10,
	}

	c1 := circle{
		r: 20,
	}

	info(s1)
	info(c1)
}
