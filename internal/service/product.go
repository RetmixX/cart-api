package service

import (
	"cart-api/internal/model/entity"
	"cart-api/internal/repository"
	"cart-api/pkg/log"
	"fmt"
)

type Producter interface {
	AllProducts() ([]entity.Product, error)
}

type ProductService struct {
	repository repository.Producter
	log        *log.Logger
}

func New(repo repository.Producter, log *log.Logger) *ProductService {
	return &ProductService{
		repository: repo,
		log:        log,
	}
}

func (p *ProductService) AllProducts() ([]entity.Product, error) {
	const op = "service.product.AllProducts"
	products, err := p.repository.FindAll()

	if err != nil {
		p.log.ErrorLog.Printf("%s: %s\n", op, err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return products, nil
}
