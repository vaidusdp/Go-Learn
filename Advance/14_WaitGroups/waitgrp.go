package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID int 
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("WorkerID %d is started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d is finished %s\n", w.ID, w.Task)
}

func main() {
	var wg sync.WaitGroup

	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i+1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	wg.Wait()
	fmt.Println("Finished")
}
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d is starting\n", id)
// 	time.Sleep(time.Second * 2)
// 	for task := range tasks {
// 		results <- task * 2
// 	}
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5
// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	for i := range numJobs {
// 		tasks <- i+1
// 	}
// 	close(tasks)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, results, &wg)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for res := range results {
// 		fmt.Println("Result:", res)
// 	}
// 	fmt.Println("End")
// }

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorker := 3

// 	wg.Add(numWorker)

// 	for i := range numWorker {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("End")
// }
