package interfaces

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type OrderService interface {
	GetAllOrders() ([]*entity.Order, error)
	GetOrdersByEmail(string) ([]*entity.Order, error)
	CreatePaymentToken(restmodel.Payment) (string, error)
	AddOrders([]restmodel.Order, string, int64) ([]*restmodel.Order, error)
	UpdateOrder(entity.Order) (*entity.Order, error)
	DeleteOrder(string, string) error
}
