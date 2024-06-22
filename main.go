package main

import (
	"go-snake/api"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	api.StartServer()
}
