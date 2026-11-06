package main

import (
	"fmt"
	"time"
)

// Go routines are just functions that leaves the main thread and runs in background and come backs to main thread when functions are finished to return any value
// Go routines dont block the execution flow and are non blocking

func main() {	
	var err error
	fmt.Println("Start")
	go sayHello() // Extracts from the main thread to a baground thread
	fmt.Println("End")

	go func() {
		err = doWork()
	}()

	go printNums()
	go printAlphabets()
	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work Completed")
	}
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello From Go")
}

func printNums() {
	for i := 0; i<5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printAlphabets() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error  {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("Error Occoured")
}