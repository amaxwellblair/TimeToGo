package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDB interface {
}

func NewRecorder(method string, handlerFunc func(w http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, "localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	h := handler{}
	handlerFunc(w, req)
	return w
}

func TestAPI_getRootHandler(t *testing.T) {
	h := handler{}
	w := NewRecorder("GET", h.getRootHandler)
	if w.Body.String() != "" {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestAPI_getMessageHandler(t *testing.T) {
	h := handler{}
	w := NewRecorder("GET", h.getMessageHandler)
}

// func TestAPI_messageHandler() {
//
// }
