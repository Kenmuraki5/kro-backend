package interfaces

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type OrderService interface {
	GetAllOrders() ([]*entity.Order, error)
	AddOrders([]restmodel.Order, restmodel.Payment) ([]*restmodel.Order, error)
	UpdateOrder(entity.Order) (*entity.Order, error)
	DeleteOrder(string, string) error
}
