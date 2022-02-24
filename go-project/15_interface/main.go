package main

import (
	"fmt"
	"math"
)

type Square struct { // class
	Length float64
}

type Circle struct { // class
	Redius float64
}

func (s *Square) Area() float64 {
	return math.Pow(s.Length, 2) // s.Length * s.Length
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Redius, 2)
}

type Shape interface {
	Area() float64
}

func sumAreas(shapes []Shape) float64 { // slices
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

func main() {
	s := &Square{20}
	fmt.Println(s.Area())

	c := &Circle{20}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	sc := sumAreas(shapes)
	fmt.Println(sc)

}
