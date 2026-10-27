package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person // Embedded Struct
	empId  string
	salary float64
}

func (p Person) introduce() {
	fmt.Printf("Hi, I'm %v and I'm %v years old\n", p.name, p.age)
}

func (e Employee) introduce() { //Overriding the introduce method
	fmt.Printf("Hi, I'm %v, employee Id: %v and I earn %.2f.\n", e.name, e.empId, e.salary)
}

func main() {
	emp := Employee{
		Person: Person{
			name: "Vaidus",
			age:  18,
		},
		empId:  "kahndg7d89",
		salary: 50000,
	}

	fmt.Println(emp.name)
	fmt.Println(emp.age)
	fmt.Println(emp.empId)
	fmt.Println(emp.salary)
	emp.introduce()
}