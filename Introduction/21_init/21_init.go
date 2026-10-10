package main

import "fmt"

func init() {
	fmt.Println("Initializing Package1...")
}

func init() {
	fmt.Println("Initializing Package2...")
}

func init() {
	fmt.Println("Initializing Package3...")
}

func main() {
	// First init runs
	// We can keep db connection inside init function as it initialises and runs b4 main function start
	fmt.Println("Inside Main")
}