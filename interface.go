package main

import (
	"math"
	"fmt"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Geomtery interface {
	area() float64
	perimeter() float64
}

func (c *Circle) area() (float64) {
	return math.Pi * c.radius * c.radius
}
func (c *Circle) perimeter() (float64) {
	return 2 * math.Pi * c.radius
}

func (r *Rectangle) area() (float64) {
	return r.height * r.width
}
func (r *Rectangle) perimeter() (float64) {
	return 2*r.width + 2*r.height
}
//method that will implement the interface struct
func measurement(g Geomtery) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
func main() {
	c := Circle{23}
	measurement(&c)
	r := Rectangle{23, 34}
	measurement(&r)
}
