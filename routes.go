package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//
// These are the URL patterns for the website
//
func LoadRoutes(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))

	s := r.PathPrefix("/community").Subrouter()
	s.HandleFunc("/{uuid}/{name}", HandleCommunity)

	r.HandleFunc("/", HomeHandler)
}
