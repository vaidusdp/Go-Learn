package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	mw "restapi/internal/api/middlewares"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
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
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Path Params
		// teachers/id
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")

		fmt.Println("Id is", userID)

		// Query Params
		// teachers/?key=value&k=v
		fmt.Println("QueryParams:", r.URL.Query())
		queryParams := r.URL.Query()
		key := queryParams.Get("key")
		k := queryParams.Get("k")
		if k == "" {
			k = "v"
		}

		fmt.Println(key, k)
	case http.MethodPost:

	case http.MethodPut:

	case http.MethodPatch:

	case http.MethodDelete:

	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:

	case http.MethodPut:

	case http.MethodPatch:

	case http.MethodDelete:

	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:

	case http.MethodPut:

	case http.MethodPatch:

	case http.MethodDelete:

	}
}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      port,
		Handler:   mw.SecurityHeaders(mw.Cors(mux)),
		// Handler:   middlewares.Cors(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
