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

// NewHandler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Settig up Routes")
	
	h.Router = mux.NewRouter()
	
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	} )

}