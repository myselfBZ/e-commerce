package routes

import (
	"e-commerce/handlers"
	"net/http"
)

func ProductRoutes(mux *http.ServeMux, h *handlers.Handler) {
	mux.HandleFunc("/products/{id}", h.ProductHandle)
	mux.HandleFunc("/products", h.GetProducts)
    mux.HandleFunc("POST /products", h.CreateProduct)
}
