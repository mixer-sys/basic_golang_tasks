package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4, Height: 6}
	fmt.Println(circle.Area())    // 78.5398...
	fmt.Println(rectangle.Area()) // 24
}
