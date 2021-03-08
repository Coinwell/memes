package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/n2n2dev/n2n2-meme/auth"
	"github.com/n2n2dev/n2n2-meme/storage"
)

func main() {

	makeCategoryCodes()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("application.go: no .env file")
	}

	initDB()
	auth.Init()
	storage.Init()
	r := initRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	fmt.Println("serving port " + port)
	http.ListenAndServe(":"+port, r)
}
