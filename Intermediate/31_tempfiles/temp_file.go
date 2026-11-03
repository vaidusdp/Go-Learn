package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	tempFile, err := os.CreateTemp("", "tempoFile")
	checkError(err)

	fmt.Println("File Name:", tempFile.Name())

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())
}