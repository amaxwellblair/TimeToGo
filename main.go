package main

import (
	"net/http"

	"github.com/amaxwellblair/TimeToGo/handlers"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	h := handlers.Handler{}
	r := mux.NewRouter()
	r.HandleFunc("/", h.GetRootHandler)
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
