package main

import (
	"e-commerce/initializers"
	"e-commerce/models"
	"log"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()

}

func main() {
    err := initializers.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Address{},&models.OrderItem{})
    if err != nil{
        log.Fatal("You messed up with migrating the database")

    }
}
