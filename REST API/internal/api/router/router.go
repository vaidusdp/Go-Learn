package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {
	// mux := http.NewServeMux()

	// mux.HandleFunc("GET /", handlers.RootHandler)
	// mux.HandleFunc("GET /execs/", handlers.ExecsHandler)
	// return mux

	tRouter := teachersRouter()
	sRouter := studentsRouter()

	tRouter.Handle("/", sRouter)
	return tRouter
}
