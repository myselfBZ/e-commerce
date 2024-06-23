package routes

import (
	"e-commerce/handlers"
	"net/http"
)

func UsersRoutes(mux *http.ServeMux, h *handlers.Handler) {
    mux.HandleFunc("/users", h.UserHandler)
}
