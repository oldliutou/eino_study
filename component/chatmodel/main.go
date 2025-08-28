package main

import (
	"log"

	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	// schema.Message()
	// schema.SystemMessage()
}
