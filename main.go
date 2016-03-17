package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	h := handler{}
	r := mux.NewRouter()
	r.HandleFunc("/", h.getRootHandler)
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
