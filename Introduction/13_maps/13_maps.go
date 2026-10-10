package main

import (
	"fmt"
	"maps"
)

func main() {

	// 1. var mapVariable m map[keyType]value
	// 2. mapVariable := make(map[keyType]valueType)
	// 3. mapVariable := map[keyType]valueType{
		// 	key1: value1,
		// 	key2: value2
		// }

	myMap := make(map[string]int)

	fmt.Println(myMap);

	myMap["A"] = 1
	myMap["B"] = 2
	myMap["D"] = 4

	fmt.Println(myMap)
	fmt.Println(myMap["B"])

	// If we write wrong key it will return default value 
	// fmt.Println(myMap["C"])

	// delete(myMap, "A") =>  deletes only key value
	// fmt.Println(myMap)

	// clear(myMap) => delets completly
	// fmt.Println(myMap)
 
	value, ok := myMap["A"]
	fmt.Println(value, ",", ok)

	myMap2 := map[string]int{
		"A": 1,
		"B": 2,
	}

	if maps.Equal(myMap, myMap2){
		fmt.Println("Same h!")
	}

	fmt.Println(myMap2)
	for k, v := range myMap{
		fmt.Println(k, ":", v)
	}

	// Nil map -> Make it first then add values
	// First make(map[String]int)
}