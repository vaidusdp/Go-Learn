package handlers

import (
	"fmt"
	"net/http"
)

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
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
