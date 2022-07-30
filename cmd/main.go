package main

import (
	"log"
	"mrLate/internal/handlers"
	"net/http"
)

func main() {
	router := handlers.NewHandler().Router
	log.Fatal(http.ListenAndServe(":8080", router))
}
