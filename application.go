package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/n2n2dev/n2n2-meme/auth"
	"github.com/n2n2dev/n2n2-meme/storage"
)

func main() {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

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

	fmt.Println("serving port " + port)
	http.ListenAndServe(":"+port, r)
}
