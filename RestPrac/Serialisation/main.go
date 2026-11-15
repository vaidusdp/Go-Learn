package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{
		Name:  "Alice",
		Email: "alice@gmail.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(jsonData))

	var unmUser User
	err = json.Unmarshal(jsonData, &unmUser)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("User:", unmUser)


	jsonDatax := `{"name": "Jhon", "email":"jhon@email.com"}`
	reader := strings.NewReader(jsonDatax)
	decoder := json.NewDecoder(reader)

	var u User
	err = decoder.Decode(&u)
	if err != nil {
		log.Fatal(err)
	} 
	fmt.Println(u)


	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	err = encoder.Encode(user)
	if err != nil {
		log.Fatal(err)
	} 
	fmt.Println(buff.String())
}