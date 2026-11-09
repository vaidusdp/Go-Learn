package main

import "fmt"

func producer(ch chan <- int){
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <- chan int, out chan <- int){
	for val := range in {
		if val % 2 == 0 {
			out <- val
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}
}

// func main() {
// 	ch := make(chan int)
// 	close(ch)
// 	close(ch)
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i 
// 		}
// 		close(ch)
// 	}()
	
// 	for val := range ch {
// 		fmt.Println(val)
// 	}
// }

// func main() {
// 	ch := make(chan int)
// 		close(ch)


// 	val, ok := <- ch
// 	if !ok{
// 		fmt.Println("Channel Is Closed")
// 		return
// 	}
// 	fmt.Println(val)
// }

// Simple Closing Channel
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		fmt.Println(val)
// 	}
// }