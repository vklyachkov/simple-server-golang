package main

import (
	"api/pkg/api"
	"log"
	"net/http"
)

func main() {
	server := api.New(":8080", http.NewServeMux())
	server.FillEndpoints()
	log.Println("Server has been started")
	log.Fatal(server.ListenAndServe())
}
