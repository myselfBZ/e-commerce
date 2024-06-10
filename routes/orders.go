package routes

import (
	"e-commerce/handlers"
	"net/http"
)

func OrdersRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/orders", handlers.OrdersHandle)
}