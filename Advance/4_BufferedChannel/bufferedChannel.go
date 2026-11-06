package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(1 * time.Second)	
// 		ch <- 1
// 		ch <- 2
// 	}()

// 	fmt.Println(<- ch)
// 	fmt.Println(<- ch)
// }

func main() { 
	// make(chan Type, capacity)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // DeadLock Capacity is full
	fmt.Println("Recieving From Buffer")
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("A:", <- ch)
		fmt.Println("Blocking Ends")
		// fmt.Println("B:", <- ch)
	}()

	fmt.Println("Blocking Starts")
	ch <- 3 
	
	fmt.Println("B:", <- ch)
	fmt.Println("C:", <- ch)

	fmt.Println("Buffered Chnnels")
}