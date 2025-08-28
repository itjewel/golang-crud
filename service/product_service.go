package service

import (
	"errors"
	"golang-crud/models"
	"golang-crud/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (pc *ProductService) GetProductService() ([]models.Product, error) {
	return pc.Repo.GetProducts()
}

func (pc *ProductService) AddProduct(p models.Product) (*models.Product, error) {
	if p.Name == "" {
		return nil, errors.New("product id wrong")
	}
	id, err := pc.Repo.Insert(p)
	if err != nil {
		return nil, err
	}
	p.ID = int(id)
	return &p, nil
}
