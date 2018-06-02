package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"time"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/{userId}", LoginHandler).Methods("PUT")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting http server.")
	log.Fatal(srv.ListenAndServe())
}
