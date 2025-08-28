package database

import (
	"fmt"
	"log"
)

func Seed() {
	// ✅ Seed Categories
	categoryQuery := `
	INSERT INTO categories (name, price) VALUES
	('Electronics', 1200.50),
	('Groceries', 300.75),
	('Books', 150.00)
	ON DUPLICATE KEY UPDATE name = VALUES(name), price = VALUES(price);
	`
	_, err := DB.Exec(categoryQuery)
	if err != nil {
		log.Fatal("Seeding categories failed:", err)
	}
	fmt.Println("✅ Categories seeding done")
	// ✅ Seed Products
	productQuery := `
	INSERT INTO products (category_id, name, stock) VALUES
	(1, 'Laptop', 10),
	(1, 'Mobile', 20),
	(2, 'Rice', 50),
	(3, 'Novel', 15)
	ON DUPLICATE KEY UPDATE name = VALUES(name), stock = VALUES(stock);
	`
	_, err = DB.Exec(productQuery)
	if err != nil {
		log.Fatal("Seeding products failed:", err)
	}
	fmt.Println("✅ Products seeding done")

	// ✅ Seed Users
	userQuery := `
	INSERT INTO users (name, email) VALUES
	('Admin', 'admin@example.com'),
	('Jewel', 'jewel@example.com')
	ON DUPLICATE KEY UPDATE name = VALUES(name), email = VALUES(email);
	`
	_, err = DB.Exec(userQuery)
	if err != nil {
		log.Fatal("Seeding users failed:", err)
	}
	fmt.Println("✅ Users seeding done")

	fmt.Println("🎉 All seeding completed successfully")
}
