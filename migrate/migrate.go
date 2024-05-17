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
	var errs = []error{
		initializers.DB.AutoMigrate(&models.Address{}),
		initializers.DB.AutoMigrate(&models.User{}),
		initializers.DB.AutoMigrate(&models.Product{}),
		initializers.DB.AutoMigrate(&models.Order{}),
		initializers.DB.AutoMigrate(&models.OrderItem{}),
	}

	for _, err := range errs {
		if err != nil {
			log.Fatal("Error Migrating", err)
			break
		}
	}
}
