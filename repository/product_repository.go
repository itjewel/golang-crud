package repository

import (
	"golang-crud/database"
	"golang-crud/models"
	"log"
)

type ProductRepository struct{}

func (p *ProductRepository) GetProducts() ([]models.Product, error) {
	res, err := database.DB.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err)
	}
	var customProducts []models.Product
	for res.Next() {
		var p models.Product
		if err := res.Scan(&p.ID, &p.Details, &p.Name, &p.Price); err != nil {
			log.Println(err)
		}
		customProducts = append(customProducts, p)
	}

	return customProducts, nil
}

func (p *ProductRepository) Insert(pr models.Product) (int64, error) {
	response, err := database.DB.Exec("INSERT INTO products (name, price, details) VALUES (?, ?,?)", pr.Name, pr.Price, pr.Details)
	if err != nil {
		return 0, nil
	}
	id, _ := response.LastInsertId()

	return id, nil
}
