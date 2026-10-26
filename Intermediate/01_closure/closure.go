package main

import "fmt"

func main() {
	// 1. Adder function is called once i = 0
	// 2. sequence is the increment function returned by adder function
	// 3. Everytime you run the sequence i++ => sequence() = adder()()
	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	// Here subtractor is the function decrements
	subtractor := func () func(int) int {
		res := 99

		return func(x int) int {
			res -= x
			return res
		}
	}()

	// Using closure subtrator
	fmt.Println(subtractor(5))
	fmt.Println(subtractor(5))
	fmt.Println(subtractor(5))
	fmt.Println(subtractor(5))
}

func adder() func() int {
	i := 0
	fmt.Println("Prev value of i")
	return func() int {
		fmt.Printf("Adding 1 to %v: \n", i)
		i++
		return i
	}
}