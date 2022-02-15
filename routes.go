package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func LoadRoutes(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))
	r.HandleFunc("/", HomeHandler)
}
