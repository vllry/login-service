package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Running")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, vars["userId"])
}