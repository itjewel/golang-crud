package database

import (
	"fmt"
	"log"
)

// Migration: Create table if not exists
func Migrate() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS categories (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			price DECIMAL(10,2) DEFAULT 0.0
		);`,

		`CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
			email VARCHAR(100) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS orders (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			category_id INT NOT NULL,
			quantity INT DEFAULT 1,
			total_price DECIMAL(10,2),
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Migration failed: %v\nQuery: %s", err, query)
		}
	}

	fmt.Println("✅ Migration complete: all tables created")
}

