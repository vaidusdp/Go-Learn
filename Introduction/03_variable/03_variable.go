package main

import "fmt"

// middleName := "Cane"
// Cant declare global variable like this

var middleName string = "Cane"

func main() {
	// 1. var age int
	// 2. var name1 string = "name1"
	// 3. var name2 = "name2"
	// 4. name3 := "name3"

	// Default Values
	// Numeric = 0, Boolean = false, String = "", Rest = nil

	// ------- Scope

	fmt.Println(middleName)
	printName()
}

func printName() {
	firstName := "Vaidus"
	fmt.Println(firstName)
}