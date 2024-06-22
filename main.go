package main

import (
	"go-snake/api"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	api.StartServer()
}
