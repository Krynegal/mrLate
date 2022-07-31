package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"mrLate/internal/configs"
	"mrLate/internal/handlers"
	"net/http"
)

func main() {
	configs.InitializeViper()
	port := viper.GetString("serverPort")
	router := handlers.NewHandler().Router
	log.Printf("Server run on port :%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
