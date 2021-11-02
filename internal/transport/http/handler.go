package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to comments service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns a pointer to Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - set up all the routes for the app
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter,
		r *http.Request) {
		fmt.Fprintf(w, "I am Alive")
	})
}
