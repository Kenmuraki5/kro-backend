package repository

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type OrderRepository interface {
	GetAllOrders() ([]*entity.Order, error)
	GetOrdersByEmail(string) ([]*entity.Order, error)
	AddOrders([]restmodel.Order, string) ([]*restmodel.Order, error)
	UpdateOrder(entity.Order) (*entity.Order, error)
	UpdateStock([]restmodel.Order) error
	DeleteOrder(string, string) error
}
