package main

import (
	"fmt"
	"net/http"
)

type database interface {
	GetMessage(int) *Message
	PutMessage(*Message) error
}

type handler struct {
	db database
}

func (h *handler) getRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}
