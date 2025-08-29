package repository

import (
	"context"
	"golang-crud/database"
	"golang-crud/models"
	"log"
)

type ProductRepository struct{}

func (p *ProductRepository) GetProducts(ctx context.Context) ([]models.Product, error) {
	res, err := database.DB.QueryContext(ctx, "SELECT * FROM products")
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

func (p *ProductRepository) Insert(ctx context.Context, pr models.Product) (int64, error) {
	response, err := database.DB.ExecContext(ctx, "INSERT INTO products (name, price, details) VALUES (?, ?,?)", pr.Name, pr.Price, pr.Details)
	if err != nil {
		return 0, nil
	}
	id, _ := response.LastInsertId()

	return id, nil
}
