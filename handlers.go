package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Running")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username := vars["userId"]
	secretPassword := r.FormValue("secretPassword")

	// TODO
	if username == "test" && secretPassword == "test" {
		token, err := generateToken(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			fmt.Println("Error generating token.")
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, token)
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println(vars["Username"])

	// INSERT user.

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, vars["userId"])
}
