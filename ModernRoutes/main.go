package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()


	mux.HandleFunc("POST /items/creat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Item Created")
	})

	mux.HandleFunc("DELETE /items/creat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Item Deleted")
	})

	mux.HandleFunc("/teachers/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Teacher ID %s", r.PathValue("id"))
	})

	mux.HandleFunc("/files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %s", r.PathValue("path"))
	})

	mux.HandleFunc("/path1/{param1}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Param1: %s", r.PathValue("param1"))	
	})

	// mux.HandleFunc("/{param2}/path2", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Param2: %s", r.PathValue("param2"))	
	// })

	http.ListenAndServe(":8080", mux)
}