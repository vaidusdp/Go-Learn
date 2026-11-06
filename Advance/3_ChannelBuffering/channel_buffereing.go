package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() { //Req only for unbuffered channel
		// ch <- 1
		time.Sleep(2 * time.Second)
		fmt.Println("2 Sec")
	}()

	go func() { //Req only for unbuffered channel
		// ch <- 1
		time.Sleep(3 * time.Second)
		fmt.Println("3 Sec")
	}()

	rcvr := <-ch
	fmt.Println(rcvr)
	fmt.Println("End")
}