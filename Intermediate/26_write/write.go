package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		return
	}

	defer file.Close()

	data := []byte("Hello World\n")
	_, err = file.Write(data)
	if err != nil {
		return
	}

	fmt.Println("Data Written Successfully")

	file, err = os.Create("WriteString.txt")
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go\n")
	if err != nil {
		return
	}
	fmt.Println("Data Written Successfully2")
}