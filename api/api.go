package api

import (
	"go-snake/api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	log.Println("Initializing the router...")
	router := mux.NewRouter()

	log.Println("Registering routes...")

	router.HandleFunc("/ws", handlers.WebsocketHandler).Methods("GET")

	addr := "0.0.0.0:8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
