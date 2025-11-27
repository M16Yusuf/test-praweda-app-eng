package main

import (
	"log"
	"test-praweda-app-eng/internal/config"
	"test-praweda-app-eng/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	// manual load ENV
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env \nCause: ", err.Error())
		// return
	}

	// initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Println("Failed to connect to DB \nCause: ", err.Error())
		return
	}
	defer db.Close()
	// else if not error
	log.Println("DB CONNECTED")

	// Inisialization engine gin
	router := router.InitRouter(db)
	router.Run("127.0.0.1:8080")
}
