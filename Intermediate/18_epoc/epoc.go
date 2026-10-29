package main

import (
	"fmt"
	"time"
)

func main() {
	// 00:00:00

	now := time.Now()
	unix:= now.Unix()
	fmt.Println(unix)

	t := time.Unix(unix, 0)
	fmt.Println(t)
}