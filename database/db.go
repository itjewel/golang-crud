package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/golang-crud?parseTime=true"



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
