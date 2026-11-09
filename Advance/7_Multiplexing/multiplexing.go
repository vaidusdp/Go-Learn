package main

import (
	"fmt"
	// "time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				// clean up activities
				return
			}
			fmt.Println("Received:", msg)
		}
	}
}

// func main() {
// 	ch := make(chan int)

// 	go func(){
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println(msg)
// 	case <- time.After(3 * time.Second):
// 		fmt.Println("Timeout")
// 	}

// }

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func(){
// 		time.Sleep(time.Second)
// 		ch1 <- 1
// 	}()

// 	go func(){
// 		time.Sleep(time.Second)
// 		ch2 <- 2
// 	}()

// 	time.Sleep(time.Second * 2)

// 	for range 2 {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println(msg1)
// 		case msg2 := <-ch2:
// 			fmt.Println(msg2)
// 		// default:
// 		// 	fmt.Println("No Channel Accepting")
// 		}
// 	}
	

// 	fmt.Println("End")
// }