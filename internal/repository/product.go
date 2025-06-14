package repository

import (
	"cart-api/internal/model/entity"
	"fmt"
	"gorm.io/gorm"
)

type Producter interface {
	FindAll() ([]entity.Product, error)
}

type Product struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Product {
	return &Product{
		db: db,
	}
}

func (p *Product) FindAll() ([]entity.Product, error) {
	const op = "repository.product.FindAll"
	var products []entity.Product

	if err := p.db.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return products, nil
}
