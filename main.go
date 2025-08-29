package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud/database"
	"golang-crud/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not loaded, using system environment variables")
	}

	database.Connect() // DB connect
	defer database.DB.Close()

	// // Reusable migration function
	// err := database.RunMigration(action)
	// if err != nil {
	// 	log.Fatal("Migration failed:", err)
	// }
	mux := http.NewServeMux()
	routes.CategoryRoutes(mux)
	routes.ProductRoutes(mux)
	routes.UserRoutes(mux)

	fmt.Println("Server running at http://localhost:8000")
	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
