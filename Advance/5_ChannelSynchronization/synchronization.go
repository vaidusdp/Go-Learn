package main

import (
	"fmt"
	"time"
)

// func main() {
// 	done := make(chan struct{})
	
// 	go func() {
// 		fmt.Println("Start")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()

// 	<- done
// 	fmt.Println("Finished")
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("Sent Value")
// 	}()

// 	time.Sleep(3 * time.Second)
// 	value := <- ch
// 	fmt.Println(value)
// }

// func main() {
// 	numGoRoutines := 3

// 	done := make(chan int, numGoRoutines)

// 	for i := range numGoRoutines {
// 		go func(id int) {
// 			fmt.Printf("Go Routine: %d\n", id)
// 			time.Sleep(time.Second)
// 			done <- id
// 		}(i)
// 	}

// 	for range numGoRoutines {
// 		<- done //Wait for all to complete for signal
// 	}

// 	fmt.Println("All Finished")
// }	

func main() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + string('0' + i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	// close(data) // Channel closed before Goroutine could send a value to the channel

	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	} // Loops over only on active channel, creates receiver each time and stops creating receiver (looping) once the channel is closed
}