package service

import "cart-api/internal/model/entity"

type ProductInter interface {
	AllProducts() ([]entity.Product, error)
}
