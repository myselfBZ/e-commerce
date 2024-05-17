package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Problem loading enviroment variables")
	}

}
