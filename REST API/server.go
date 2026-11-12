package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Server")
	})

	const port string = ":3000"

	fmt.Println("Started Server")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		os.Exit(1)
	}

	
}