package service

import "cart-api/internal/model/entity"

type OrderInter interface {
	AllOrders() ([]entity.Order, error)
	CreateOrder() error
	ChangeStatus(newStatus *entity.Status) error
}
