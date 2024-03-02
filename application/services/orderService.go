// services/game_service.go

package services

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	omise "github.com/Kenmuraki5/kro-backend.git/pkg/payment"
	"github.com/google/uuid"
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

func (s *OrderService) AddOrders(order []restmodel.Order, payment restmodel.Payment) ([]*restmodel.Order, error) {
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

	err := s.orderRepository.UpdateStock(gameOrders, consoleOrders)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	client, err := omise.GetOmiseClient()
	if err != nil {
		return nil, fmt.Errorf("error creating Omise client: %v", err)
	}
	token, err := omise.CreateToken(client, payment)
	if err != nil {
		return nil, fmt.Errorf("error creating token: %v", err)
	}
	orderId := uuid.NewString()
	err = omise.CreateChargeByToken(client, token, orderId, payment.Total)
	if err != nil {
		return nil, fmt.Errorf("error creating charge by token: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error creating Omise client: %v", err)
	}

	addedOrder, err := s.orderRepository.AddOrders(order, orderId)
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
