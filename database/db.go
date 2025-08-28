package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/golang-migrate/migrate/v4/database/mysql"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	// driver := os.Getenv("DB_DRIVER")
	// appmode := os.Getenv("APP_MODE")

	//dsn := "root:admin@tcp(127.0.0.1:3306)/go-gin-crud"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)

	// Open connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB Connection Error:", err)
	}

	// Ping for check
	err = db.Ping()
	if err != nil {
		log.Fatal("DB Ping Error:", err)
	}

	fmt.Println(" Successfully connected to MySQL")
	DB = db
}
