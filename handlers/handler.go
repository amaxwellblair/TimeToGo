package handlers

import (
	"fmt"
	"net/http"

	"github.com/amaxwellblair/slackdown/models"
)

// Handler will be a wrapper for all handler functions
type Handler struct {
	Store models.Store
}

// GetRootHandler handles requests to the root
func (h *Handler) GetRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}
