package main

import (
	"log"

	"github.com/PatrickELee/Via/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error with .env")
	}

	vs := server.CreateViaWebServers()
	vs.ListenAndServe()
}
