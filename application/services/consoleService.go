package services

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type ConsoleService struct {
	consoleRepository repository.ConsoleRepository
}

func NewConsoleService(
	consoleRepository repository.ConsoleRepository,
) *ConsoleService {
	return &ConsoleService{consoleRepository: consoleRepository}
}

func (s *ConsoleService) GetAllConsoles() ([]*entity.Console, error) {
	return s.consoleRepository.GetAllConsoles()
}

func (s *ConsoleService) AddConsole(console restmodel.Console) (*restmodel.Console, error) {
	addedConsole, err := s.consoleRepository.AddConsole(console)
	if err != nil {
		return nil, err
	}

	return addedConsole, nil
}

func (s *ConsoleService) UpdateConsole(console entity.Console) (*entity.Console, error) {
	updatedConsole, err := s.consoleRepository.UpdateConsole(console)
	if err != nil {
		return nil, err
	}
	return updatedConsole, nil
}

func (s *ConsoleService) ReleaseStockConsole(order entity.Order) error {
	err := s.consoleRepository.ReleaseStockConsole(order)
	if err != nil {
		return fmt.Errorf("failed to Release Stock game ID %s: %w", order.ProductId, err)
	}

	return err
}

func (s *ConsoleService) DeleteConsole(id string) error {
	err := s.consoleRepository.DeleteConsole(id)
	if err != nil {
		return fmt.Errorf("failed to delete console with ID %s: %w", id, err)
	}

	return err
}
