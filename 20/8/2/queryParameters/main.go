package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "category: %s\n", q["category"][0])
	fmt.Fprintf(w, "id: %s\n", q["id"][0])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", QueryHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
