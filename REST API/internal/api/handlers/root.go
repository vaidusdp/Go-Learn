package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Func")
	w.Write([]byte("Welcome To School API"))
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
}
