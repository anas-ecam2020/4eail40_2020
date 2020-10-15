package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	String() string
}

type Rectangle struct {
	width  float64
	length float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Square of width %v and length %v and area: %v", r.width, r.length, r.Area())
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %v and area: %v", c.radius, c.Area())
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}
}
