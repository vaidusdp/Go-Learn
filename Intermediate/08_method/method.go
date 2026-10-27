package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
} 

type MyInt int

func (m *MyInt) Define(x int) int {
	return x
}

func main() {
	rect := Rectangle{
		length: 10,
		width:  9,
	}

	area := rect.Area()
	fmt.Println("Area:", area)

	rect.Scale(10)

	area = rect.Area()
	fmt.Println("Area:", area)

	num := MyInt(-5)
	fmt.Println(num)
	newNum := num.Define(2)
	fmt.Println(newNum)
}
