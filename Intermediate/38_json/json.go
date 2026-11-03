package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age,omitempty"`
	Email string `json:"email"`
}

type Employee struct {
	Name  string `json:"name"`
	EmpId   int    `json:"emp_id,omitempty"`
}

func main() {
	p := Person{Name: "Vaidus", Email: "vaidus2608m@gmail.com"}

	jsonData, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	data := `{"name": "Jenny", "emp_id": "0009"}`

	var emp Employee
	json.Unmarshal([]byte(data), &emp)

	fmt.Println(emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.EmpId)

	data2 := `{"name": "Vexor", "age": 30}`
	var datam map[string]interface{}
	json.Unmarshal([]byte(data2), &datam)

	fmt.Println(datam)
	fmt.Println(datam["age"])
}
