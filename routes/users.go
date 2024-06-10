package routes

import (
	"e-commerce/handlers"
	"net/http"
)

func UsersRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", handlers.UserHandler)
}
