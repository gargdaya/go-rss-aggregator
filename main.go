package main

import (
	"os"

	"github.com/joho/godotenv"
)


func main() {
	println("Hello, World!")

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	println("Listening on port " +  port)
}