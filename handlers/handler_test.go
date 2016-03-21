package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amaxwellblair/TimeToGo/handlers"
)

type mockDB interface {
}

func NewRecorder(method string, handlerFunc func(w http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, "localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	handlerFunc(w, req)
	return w
}

func TestAPI_GetRootHandler(t *testing.T) {
	h := handlers.Handler{}
	w := NewRecorder("GET", h.GetRootHandler)
	if w.Body.String() != "" {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestAPI_getMessageHandler(t *testing.T) {
	// h := Handler{}
	// w := NewRecorder("GET", h.getMessageHandler)
}
