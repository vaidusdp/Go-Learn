package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"-"` // - means omit everytime
}

func main() {
	p := Person{Name: "Vexor", Age: 30}

	jsonData, _ := json.Marshal(p)

	fmt.Println(string(jsonData))

}