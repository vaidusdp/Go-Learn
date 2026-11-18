package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	// "time"

	mw "restapi/internal/api/middlewares"
)

type Teacher struct {
	ID        int
	FirstName string
	LastName  string
	Subject   string
	Class     string
}

var (
	teachers = make(map[int]Teacher)
	mutex    = &sync.Mutex{}
	nextId   = 1
)

func init() {
	teachers[nextId] = Teacher{
		nextId,
		"Jhon",
		"Doe",
		"CN",
		"2-A",
	}

	nextId++
	teachers[nextId] = Teacher{
		nextId,
		"Jane",
		"Smith",
		"Spring",
		"1-A",
	}

	nextId++
	teachers[nextId] = Teacher{
		nextId,
		"Jane",
		"Doe",
		"Spring",
		"1-A",
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")
		teacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		} 
		response := struct{
			Status string `json:"status"`
			Count int `json:"count"`
			Data []Teacher `json:"data"`
		}{
			Status: "success",
			Count: len(teachers),
			Data: teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalln("Unable to convert:", err)
	}

	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

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
		getTeachersHandler(w, r)
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
		fmt.Println("Query:", r.URL.Query())
		fmt.Println("Name:", r.URL.Query().Get("name"))

		err := r.ParseForm()
		if err != nil {
			return
		}
		fmt.Println("Form:", r.Form)
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

	// rl := mw.NewRateLimiter(5, time.Minute)
	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(rl.Middleware(mw.ResoponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	// secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.SecurityHeaders, mw.ResoponseTimeMiddleware, rl.Middleware, mw.Cors)
	secureMux := mw.Cors(mux)

	server := &http.Server{
		Addr:    port,
		Handler: secureMux,
		// Handler:   middlewares.Cors(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}

type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
