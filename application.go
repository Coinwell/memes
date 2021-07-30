package main

import (
	"log"
	"net/http"
	"os"

	"github.com/getzion/memes/auth"
	"github.com/getzion/memes/storage"
)

func main() {
	makeCategoryCodes()

	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("application.go: no .env file")
	// }

	initDB()
	auth.Init()
	storage.Init()
	r := initRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Println("serving port " + port)
	http.ListenAndServe(":"+port, r)
}
