package main

import "fmt"

func main() {
	ch := make(chan int)
	producer(ch)	
	consumer(ch)
}

// Recieve only
func producer(ch chan <- int){
	go func(ch chan <- int){
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}(ch)
}

// Send only
func consumer(ch <- chan int){
	for val := range ch {
		fmt.Println(val)
	}
}