package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Handling Incoming Orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Handling Users")
	})

	port := 3000

	fmt.Println("Server is running on port: 3000")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {

	}
}