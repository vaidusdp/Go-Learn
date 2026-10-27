package main

import "fmt"

type Geometry interface {
	// Interface is like a set of function that a specific type must implement to be that interface
	// Can have additional method
	
	area() float64
	perim() float64
}

type Rect struct {
	length, width float64
}

func (r Rect) area() float64 {
	return r.length * r.width
}

func (r Rect) perim() float64 {
	return 2 * (r.length + r.width)
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := Rect{
		length: 10,
		width:  5,
	}

	measure(r)
}