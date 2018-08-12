package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	s.router.Handle("/metrics", promhttp.Handler())
	s.router.Handle("/test", AddContext(http.HandlerFunc(RegisterHandler)))
	s.router.HandleFunc("/register/{userId}", RegisterHandler).Methods("POST")
	s.router.HandleFunc("/login/{userId}", LoginHandler).Methods("POST")

	http.Handle("/", s.router)

	return &s
}

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		ctx := context.WithValue(r.Context(), "Username", "testval")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *server) start() {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting http server.")
	log.Fatal(srv.ListenAndServe())
}
