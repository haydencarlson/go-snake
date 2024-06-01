package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func GameHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GameHandler was hit")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to the Game!</h1>")
}
