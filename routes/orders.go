package routes

import (
	"e-commerce/handlers"
	"net/http"
)

func OrdersRoutes(mux *http.ServeMux, h *handlers.Handler) {
    mux.HandleFunc("/orders/{id}", h.OrdersHandle)
    mux.HandleFunc("/orders", h.GetOrders)
}
