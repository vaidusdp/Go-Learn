package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	msg1 := "Hello \nGo" 	// DoubleQoutes -> Escape Sequence
	msg2 := "Hello \tGo"
	msg3 := "Hello \rGoxy"
	rawMsg := `Hello \nGo` // BackTick -> Raw String 

	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
	fmt.Println(rawMsg)


	// Strings are array of unicode chars
	fmt.Println(len(msg1))
	fmt.Println(len(msg2))
	fmt.Println(len(msg3))
	fmt.Println(len(rawMsg))

	//ZeroIndexing
	fmt.Println(msg1[0]) // ASCII Value
	fmt.Println("Hello" + " World")

	str1 := "Apple"
	str2 := "Banana"
	fmt.Println(str1 < str2) 
	// Lexographical Comparison
	// Comparing ASCII value
	// According to dictionary

	// Iterate
	for i, c := range msg1 {
		fmt.Printf("Index %v is %v\n", i, c) //Format verbs
	}

	// RuneCount
	fmt.Println("Rune Count:", utf8.RuneCountInString(msg1))

	// String are immutable

	
	var ch rune = 'a'
	fmt.Println(ch)
	fmt.Printf("%c\n", ch)

	cstr := string(ch)
	fmt.Printf("%T\n", ch)
	fmt.Printf("%T\n", cstr)
}