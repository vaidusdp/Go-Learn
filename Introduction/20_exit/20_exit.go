package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Defer")
	fmt.Println("Starting")
	os.Exit(1)
	// After exit everything stops completly
	fmt.Println("End")
}