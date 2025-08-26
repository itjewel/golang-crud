package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud/database"
	"golang-crud/routes"
)

func main() {
	database.Connect()

	mux := http.NewServeMux()
	routes.CategoryRoutes(mux)

	fmt.Println("Server running at http://localhost:8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
