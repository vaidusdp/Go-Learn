package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: No is Negetive")
	}

	return 2, nil
}

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eProcess() error {
	return &myError{"Custom Error Message"}
}

func main() {
	res1, err1 := sqrt(16)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(res1)
	}

	res2, err2 := sqrt(-16)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res2)
	}

	err3 := eProcess()
	if err3 != nil {
		fmt.Println(err3)
	}
}