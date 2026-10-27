package main

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
	err error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s\n%v", e.code, e.message, e.err)
}

// func do() error {
// 	return &customError{
// 		code: 500,
// 		message: "Something went wrong!",
// 	}
// }

func do() error {
	e := doelse()
	if e != nil {
		return &customError{
			code: 500,
			message: "Something went wrong!",
			err: e,
		}
	}

	return nil
}

func doelse() error {
	return errors.New("Internal Error")
}

func main() {
	err := do()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Op completed")
}
