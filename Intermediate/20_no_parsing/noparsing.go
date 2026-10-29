package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "123456"
	num, _ := strconv.Atoi(numStr)
	fmt.Println(num+1)

	numistr, _ := strconv.ParseInt(numStr, 10, 32)
	fmt.Println(numistr)

	floatstr := "3.14"
	floatval, _ := strconv.ParseFloat(floatstr, 64)
	fmt.Println(floatval)

	bin := "101010"
	dec, _ := strconv.ParseInt(bin, 2, 64)
	fmt.Println(dec)
}