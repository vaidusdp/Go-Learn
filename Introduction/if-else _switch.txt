package main

import "fmt"

func main() {
	// if cond {
	// 	block of code
	// }

	// switch expression{
	// 	case value1:
	// 		code
	// 	case value2:
	// 		code
	// 	default:
	// 		code
	// }

	// No need to mention break at the end of case like other languages
	// Fall Through => Checks Next Case

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("Apple Lelo")
		fallthrough
	case "banana":
		fmt.Println("Banana Lelo")
	default:
		fmt.Println("Jo Lena h Lelo")
	}

}