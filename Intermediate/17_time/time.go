package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(time.Now())

	// specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
	// fmt.Println(specificTime)

	// parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") // Mon Jan 2 15:04:05 MST 2006
	// parsedTime1, _ := time.Parse("06-01-02", "20-05-01") // Mon Jan 2 15:04:05 MST 2006
	// fmt.Println(parsedTime)
	// fmt.Println(parsedTime1)


	t := time.Now()
	// fmt.Println(t.Format("06-01-02 04:15"))

	// oneDayLater := t.Add(time.Hour * 24)
	// fmt.Println(oneDayLater.Weekday())
	// fmt.Println(t.Round(time.Hour))

	fmt.Println(t.Truncate(time.Hour))
}