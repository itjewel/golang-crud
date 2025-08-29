package main

import (
	"fmt"
	"golang-crud/database"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run database/migrate_cli.go [up|down]")
		return
	}

	action := os.Args[1]

	database.Connect()
	defer database.DB.Close()

	if err := database.RunMigration(action); err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Migration completed successfully!")
}
