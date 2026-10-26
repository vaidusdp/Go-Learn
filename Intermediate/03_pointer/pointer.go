package main

import "fmt"

func main() {
	// Pointer store memory address of another variable
	// Easy to modify and pass in function

	var ptr *int

	var a int = 10
	ptr = &a // refrencing

	fmt.Println(a)
	fmt.Println(ptr) // location

	// ** -> derefrencing a variable
	fmt.Println(*ptr) // derefrencing
	incrementPointer(ptr)
	fmt.Println(*ptr) // derefrencing
}

func incrementPointer(ptr *int) (p *int) {
	p = ptr
	*p++
	return
}