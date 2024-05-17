package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	credentials := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(credentials), &gorm.Config{})

	if err != nil {
		log.Fatal("Error opening database.")
	}
}
