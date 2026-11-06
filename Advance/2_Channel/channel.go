package main

import (
	"fmt"
	"time"
)

func main() {
	// name := make(chan type)
	greeting := make(chan string)
	// greetString := "Hello"

	go func() {
		// greeting <- greetString //Was Blocking
		// greeting <- "World" 

		for _,e := range "abcde" {
			greeting <- string(e)
		}
	}()

	// reciever := <-greeting
	// fmt.Println(reciever)
	// reciever = <-greeting
	// fmt.Println(reciever)

	// go func() {
	// 	reciever := <-greeting
	// 	fmt.Println(reciever)
	// 	reciever = <-greeting
	// 	fmt.Println(reciever)
	// }()

	for range 5 {
		recvr := <- greeting
		fmt.Println(recvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End")
}