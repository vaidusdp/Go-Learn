package main

// Format
import "fmt"

// Every go program must have main function as this is the entry point
func main() {
	fmt.Println("Hello World!")
}

// go run hello.go 
// go build hello.go -> Maked a binary executable file 
// If the go file is changed even after making executable the executable wont change.

// Auto Garbage Collection
// C/C++ -> COMPILE -> EXEC -> RUNS WITHOUT ENVIORNMENT
// GO -> RUNTIME -> MACHINE CODE -> RUNS