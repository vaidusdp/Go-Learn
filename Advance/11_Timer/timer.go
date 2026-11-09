package main

import (
	"fmt"
	"time"
)

func main(){
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	for range 2 {
		select {
		case <- timer1.C:
			fmt.Println("Expired-1")
		case <- timer2.C:
			fmt.Println("Expired-2")
		}	
	}
	
}

// func main() {
// 	timer := time.NewTimer(2 * time.Second)
// 	go func(){
// 		<- timer.C
// 		fmt.Println("Delayed Executed")
// 	}()

// 	fmt.Println("Waiting...")
// 	time.Sleep(time.Second * 3)
// 	fmt.Println("End")
// }

// func longRunningOp() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(5 * time.Second)
// 	done := make(chan bool)

// 	go func(){
// 		longRunningOp()
// 		done <- true
// 	}()

// 	select{
// 	case <- timeout:
// 		fmt.Println("Timed Out")
// 	case <- done:
// 		fmt.Println("Op Completed")
// 	}
// }

// func main() {
// 	fmt.Println("Start")
// 	timer := time.NewTimer(2 * time.Second)
// 	fmt.Println("Waiting for value")
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Stopped")
// 	}
// 	fmt.Println("Reset")
// 	timer.Reset(time.Second)
// 	<- timer.C
// 	fmt.Println("Expired")
// }