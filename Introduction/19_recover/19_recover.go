package main

import "fmt"

func main(){
	process()
	fmt.Println("Returned from process")
}

func process() {
	// Used to capture panic and run code smoothly
	defer func() {
		if r := recover(); r!=nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process...")
	panic("Something went wrong!!")
	fmt.Println("End Process") 
}