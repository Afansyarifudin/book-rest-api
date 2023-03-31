package main

import (
	"book-rest-api/app"
	"book-rest-api/config"

	"github.com/joho/godotenv"
)

func init() {
	// init load os
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// inisialisasi postgres

	err = config.InitPostgres()
	if err != nil {
		panic(err)
	}

}

func main() {
	app.StartApplication()
}
