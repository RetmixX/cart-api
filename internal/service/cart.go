package service

import "cart-api/internal/model/entity"

type CartInter interface {
	ViewCart() (*entity.Cart, error)
	AddProductInCart(product *entity.Product) error
	RemoveProductFromCart(idProduct int) error
}
