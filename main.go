package main

import (
	"e-commerce/initializers"
	"e-commerce/routes"
	"log"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

var mux *http.ServeMux

func main() {

	var port = os.Getenv("PORT")
	routes.OrdersRoutes(mux)
	routes.ProductRoutes(mux)
	routes.UsersRoutes(mux)
	var server = http.Server{
		Handler: mux,
		Addr:    port,
	}
	log.Println("Server is running on http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
