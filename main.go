package main

import "fmt"

type Shape interface {
	Area() int64
}

type Feature struct {
	color    string
	material string
}

type Rectangle struct {
	Feature
	width, height int64
}

func NewRectangle(color, material string, width, height int64) *Rectangle {
	return &Rectangle{
		Feature: Feature{color: color, material: material},
		width:   width,
		height:  height,
	}
}

func (r *Rectangle) Area() int64 {
	return r.width * r.height
}

type Circle struct {
	Feature
	radius int64
}

func NewCircle(color, material string, radius int64) *Circle {
	return &Circle{
		Feature: Feature{color: color, material: material},
		radius:  radius,
	}
}

func (c *Circle) Area() int64 {
	return c.radius * c.radius
}

func main() {
	r := NewRectangle("red", "wood", 10, 5)
	c := NewCircle("blue", "metal", 5)
	s := []Shape{r, c}

	for _, shape := range s {
		fmt.Printf("%+v\n", shape)
		fmt.Println(shape.Area())
	}
}
