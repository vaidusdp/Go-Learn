package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/models"
)

var (
	teachers = make(map[int]models.Teacher)
	nextId   = 1
)

func init() {
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "Jhon",
		LastName:  "Doe",
		Subject:   "CN",
		Class:     "2-A",
	}

	nextId++
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "Smith",
		Subject:   "Spring",
		Class:     "1-A",
	}

	nextId++
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "Doe",
		Subject:   "Spring",
		Class:     "1-A",
	}

	nextId++
}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

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
	router := router.Router()
	secureMux := mw.Cors(router)

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
