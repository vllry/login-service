package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	s := server{
		mux.NewRouter(),
	}

	s.router.HandleFunc("/", IndexHandler)
	s.router.HandleFunc("/{userId}", LoginHandler).Methods("PUT")

	http.Handle("/", s.router)

	return &s
}

func (s *server) start() {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting http server.")
	log.Fatal(srv.ListenAndServe())
}
