package main

import (
	"flag"
	"fmt"
	// "os"
)

func main() {
	// fmt.Println("Command:", os.Args[0])

	// for i, arg := range os.Args {
	// 	fmt.Println("Argument", i, ":", arg)
	// }

	var name string
	var age int 

	flag.StringVar(&name, "name", "Vaidus", "Name of the User")
	flag.IntVar(&age, "age", 18, "Age of the User")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}