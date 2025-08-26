package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	// Correct DSN: user:password@tcp(host:port)/dbname
	dsn := "root:admin@tcp(127.0.0.1:3306)/go-gin-crud"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB Connection Error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB Ping Error:", err)
	}

	fmt.Println("✅ Successfully connected to MySQL")
	DB = db
}
