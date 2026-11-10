package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoRoutines := 5
	wg.Add(numGoRoutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGoRoutines {
		go increment()
	}

	wg.Wait()
	fmt.Println("Final Value:", counter)
}

// type Counter struct {
// 	mu sync.Mutex
// 	value int 
// }

// func (c *Counter) increment() {
// 	c.mu.Lock()
// 	c.value++
// 	defer c.mu.Unlock()
// } 
// func (c *Counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.value
// } 

// func main() {
// 	var wg sync.WaitGroup
// 	counter := &Counter{}

// 	numGoRoutine := 10
// 	wg.Add(numGoRoutine)

// 	for range numGoRoutine {
// 		// wg.Add(1)
// 		go func(){
// 			defer wg.Done()
// 			for range 1000 {
// 				// counter.value++
// 				counter.increment()
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Final Value:", counter.getValue())
// }