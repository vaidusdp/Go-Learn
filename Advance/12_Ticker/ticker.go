package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(4 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println(tick)
		case <-stop:
			fmt.Println("Stopping")
			return
		}
	}
}

// func periodTask() {
// 	fmt.Println(time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for range 5 {
// 		select{
// 		case <- ticker.C:
// 			periodTask()
// 		}
// 	}
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	i := 1
// 	for range 5 {
// 		fmt.Println(i)
// 		i*=2
// 	}
// }
