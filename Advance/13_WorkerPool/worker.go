package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personId   int
	numTickets int
	cost       int
}

func ticketProcessor(requests <-chan ticketRequest, result chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket of personId %d with cost %d\n", req.numTickets, req.personId, req.cost)
		time.Sleep(time.Second)
		result <- req.personId
	}
}

func main() {
	numReq := 5
	price := 5

	ticketReq := make(chan ticketRequest, numReq)
	ticketRes := make(chan int)

	for range 3 {
		go ticketProcessor(ticketReq, ticketRes)
	}

	for i := range numReq {
		ticketReq <-ticketRequest{personId: i+1, numTickets: (i+1) * 2, cost: (i+1) * price}
	}
	close(ticketReq)

	for range numReq {
		fmt.Printf("Ticket for personId %d processed successfully\n", <-ticketRes)
	}
}

// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Println(id, ":", task)
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }

// func main() {
// 	noWorker := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	for i := range noWorker {
// 		go worker(i, tasks, results)
// 	}

// 	for i := range numJobs {
// 		tasks <- i
// 	}

// 	close(tasks)

// 	for range results {
// 		res := <-results
// 		fmt.Println("Res:", res)
// 	}
// }
