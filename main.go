package main

import (
	"e-commerce/handlers"
	"e-commerce/initializers"
	"log"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	var mux = http.NewServeMux()
	var port = os.Getenv("PORT")
	mux.HandleFunc("GET /products/{id}", handlers.GetProduct)
	mux.HandleFunc("POST /products", handlers.CreateProduct)
	mux.HandleFunc("GET /products", handlers.GetProducts)
	mux.HandleFunc("DELETE /products/{id}", handlers.DeleteProduct)
	mux.HandleFunc("PUT /products/{id}", handlers.UpdateProduct)
	mux.HandleFunc("POST /users", handlers.CreateUser)
	mux.HandleFunc("GET /users", handlers.GetUsers)
	mux.HandleFunc("POST /orders", handlers.CreateOrder)
	var server = http.Server{
		Handler: mux,
		Addr:    port,
	}
	log.Println("Server is running on http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
