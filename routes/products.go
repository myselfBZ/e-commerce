package routes

import (
	"e-commerce/handlers"
	"net/http"
)

func ProductRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/products/{id}", handlers.ProductHandle)
	mux.HandleFunc("/products", handlers.GetProducts)
}
