package e

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return (r.Height) * (r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return (c.Radius) * (c.Radius) * math.Pi
}

func E1print() {
	r1 := Rectangle{Width: 12, Height: 10}
	r2 := Rectangle{Width: 9, Height: 4}
	fmt.Println("Area of r1 is:", r1.Area())
	fmt.Println("Area of r2 is:", r2.Area())

	c1 := Circle{Radius: 12}
	c2 := Circle{Radius: 9}
	fmt.Println("Area of c1 is:", c1.Area())
	fmt.Println("Area of c2 is:", c2.Area())
}
