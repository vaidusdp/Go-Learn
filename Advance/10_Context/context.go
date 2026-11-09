package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context){
	for {
		select {
		case <- ctx.Done():
			fmt.Println("Cancelled: ", ctx.Err())
		default:
			fmt.Println("Working")
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ctx := context.Background()
	ctx,cancel := context.WithTimeout(ctx, 2 * time.Second) //Timer Start

	go func() {
		time.Sleep(2 * time.Second)
		defer cancel()
	}()

	ctx = context.WithValue(ctx, "reqId", "allakdjfh")
	go doWork(ctx)

	time.Sleep(5 * time.Second)

	reqId := ctx.Value("reqId")

	if reqId != nil {
		fmt.Println("ReqId:", reqId)
	} else {
		fmt.Println("No ReqId")
	}
}

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select{
// 	case <- ctx.Done():
// 		return "OP Cancelled"
// 	default:
// 		if num % 2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO()

// 	res := checkEvenOdd(ctx, 5)
// 	fmt.Println(res)

// 	ctx = context.Background()
// 	ctx,cancel := context.WithTimeout(ctx, 2 *  time.Second)
// 	defer cancel()

// 	res = checkEvenOdd(ctx, 10)
// 	fmt.Println(res)

// 	time.Sleep(3 * time.Second)
// 	res = checkEvenOdd(ctx, 13)
// 	fmt.Println(res)
// }

// func main() {
// 	todoContext := context.TODO()
// 	contextBkg := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "jhon")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	bkg := context.WithValue(contextBkg, "city", "New York")
// 	fmt.Println(bkg)
// 	fmt.Println(bkg.Value("city"))
// }