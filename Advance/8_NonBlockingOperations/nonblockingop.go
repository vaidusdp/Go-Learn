package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)

	// select {
	// case msg := <-ch:
	// 	fmt.Println(msg)
	// default:
	// 	fmt.Println("No msg")
	// }

	// select{
	// case ch <- 1:
	// 	fmt.Println("Sent Message")
	// default:
	// 	fmt.Println("CHannel is not ready")
	// }

	data := make(chan int)
	quit := make(chan bool)

	go func(){
		for {
			select{
			case d := <- data:
				fmt.Println("Data Recievved:", d)
			case <- quit:
				fmt.Println("Stoping...")
			default:
			fmt.Println("Waiting for data")
			time.Sleep(500*time.Millisecond)
			}
		}
	}()
	
	for i := range 5{
		data <- i
		time.Sleep(1000*time.Millisecond)
	}
	quit <- true
}