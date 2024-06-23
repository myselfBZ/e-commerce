package main

import (
	"e-commerce/handlers"
	"e-commerce/initializers"
	"e-commerce/models"
	"e-commerce/routes"
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
    var handler = handlers.NewHandler(&models.Product{}, &models.User{}, &models.Order{}, &models.OrderItem{})
	routes.OrdersRoutes(mux, handler)
	routes.ProductRoutes(mux, handler)
	routes.UsersRoutes(mux, handler)
	var server = http.Server{
		Handler: mux,
		Addr:    port,
	}
	log.Println("Server is running on http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
