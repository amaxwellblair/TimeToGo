package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Handler will be a wrapper for all handler functions
type Handler struct {
	db *sql.DB
}

func (h *Handler) GetRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}
