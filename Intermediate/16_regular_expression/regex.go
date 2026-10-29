package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said, \"I am great\"")

	//Compile Regex for email

	reg := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	email1 := "user@email.com"
	email2 := "useremail.com"

	fmt.Println("Email1:", reg.MatchString(email1))
	fmt.Println("Email2:", reg.MatchString(email2))

	reg = regexp.MustCompile(`(\d){4}-(\d){2}-(\d){2}`)

	date := "2026-03-26"
	submatches := reg.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// i - case_insesitive
	// m - multiline_model
	// s - do_not_matches
}