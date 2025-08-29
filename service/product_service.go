package service

import (
	"context"
	"errors"
	"golang-crud/models"
	"golang-crud/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (pc *ProductService) GetProductService(ctx context.Context) ([]models.Product, error) {
	return pc.Repo.GetProducts(ctx)
}

func (pc *ProductService) AddProduct(ctx context.Context, p models.Product) (*models.Product, error) {
	if p.Name == "" {
		return nil, errors.New("product id wrong")
	}
	id, err := pc.Repo.Insert(ctx, p)
	if err != nil {
		return nil, err
	}
	p.ID = int(id)
	return &p, nil
}
