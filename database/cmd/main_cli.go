package main

import (
	"golang-crud/database"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("no match file")
		return
	}

	database.Connect()
	filename := os.Args[1]
	// log.Println(command)
	database.RunMigration(filename)
}
