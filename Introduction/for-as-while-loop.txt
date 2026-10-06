package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Guessing Game:")
	fmt.Println("I choosed the no between 1 and 100")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Woo Hoo!")
			break
		} else if guess < target {
			fmt.Println("Go Bigger")
		} else if guess > target {
			fmt.Println("Go Smaller")
		}
	}
}