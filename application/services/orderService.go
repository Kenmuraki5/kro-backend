// services/game_service.go

package services

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type OrderService struct {
	orderRepository   repository.OrderRepository
	gameRepository    repository.GameRepository
	consoleRepository repository.ConsoleRepository
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	gameRepository repository.GameRepository,
	consoleRepository repository.ConsoleRepository,
) *OrderService {
	return &OrderService{
		orderRepository:   orderRepository,
		gameRepository:    gameRepository,
		consoleRepository: consoleRepository,
	}
}

func (s *OrderService) GetAllOrders() ([]*entity.Order, error) {
	return s.orderRepository.GetAllOrders()
}

func (s *OrderService) AddOrders(order []restmodel.Order) ([]*restmodel.Order, error) {
	gameOrders := make([]restmodel.Order, 0)
	consoleOrders := make([]restmodel.Order, 0)

	for _, item := range order {
		switch item.Type {
		case "Game":
			gameOrders = append(gameOrders, item)
		case "Console":
			consoleOrders = append(consoleOrders, item)
		default:
			fmt.Printf("Unsupported order type: %s\n", item.Type)
		}
	}

	if len(gameOrders) > 0 {
		err := s.gameRepository.UpdateStockGame(gameOrders)
		if err != nil {
			return nil, err
		}
	}

	if len(consoleOrders) > 0 {
		err := s.consoleRepository.UpdateStockConsole(consoleOrders)
		if err != nil {
			return nil, err
		}
	}
	addedOrder, err := s.orderRepository.AddOrders(order)
	if err != nil {
		return nil, err
	}

	return addedOrder, nil
}

func (s *OrderService) UpdateOrder(order entity.Order) (*entity.Order, error) {
	updatedOrder, err := s.orderRepository.UpdateOrder(order)
	if err != nil {
		return nil, err
	}

	if order.Status == "Cancel" {
		switch order.Type {
		case "Game":
			err = s.gameRepository.ReleaseStockGame(order)
			if err != nil {
				return nil, err
			}
		case "Console":
			err = s.consoleRepository.ReleaseStockConsole(order)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown order type: %s", order.Type)
		}
	}

	return updatedOrder, nil
}

func (s *OrderService) DeleteOrder(orderId, productId string) error {
	err := s.orderRepository.DeleteOrder(orderId, productId)
	if err != nil {
		return fmt.Errorf("failed to delete game with orderId %s, productId %s: %w", orderId, productId, err)
	}

	return err
}
