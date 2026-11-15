package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// fmt.Fprintf(w, "Hello Root Func")
		w.Write([]byte("Hello Root Route"))
		fmt.Println(r.Method)

		// x-www-form-urlencoded
		// err := r.ParseForm()
		// if err != nil {
		// 	http.Error(w, "Error Parsing Form", http.StatusBadRequest)
		// 	return
		// }

		// fmt.Println("Form:", r.Form)

		// response := make(map[string]interface{})
		// for key, value := range r.Form {
		// 	response[key] = value[0]
		// }
		// fmt.Println("Sadda Response:", response)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		defer r.Body.Close()
		
		fmt.Println(string(body))
	})

	http.HandleFunc("/teacher", func(w http.ResponseWriter, r *http.Request){
		// fmt.Fprintf(w, "Hello Root Func")
		w.Write([]byte("Hello Teachers Route"))
	})

	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request){
		// fmt.Fprintf(w, "Hello Root Func")
		w.Write([]byte("Hello Students Route"))
	})

	http.HandleFunc("/exec", func(w http.ResponseWriter, r *http.Request){
		// fmt.Fprintf(w, "Hello Root Func")
		w.Write([]byte("Hello Execs Route"))
	})

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}