package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// basic GET
	mux.HandleFunc("GET /path/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})

	// path parmams with wildcard {id}
	mux.HandleFunc("/task/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})

	fmt.Println("Server listening on port 8090")
	http.ListenAndServe("localhost:8090", mux)
}
