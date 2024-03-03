// services/game_service.go

package services

import (
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type GameService struct {
	gameRepository repository.GameRepository
}

func NewGameService(
	gameRepository repository.GameRepository,
) *GameService {
	return &GameService{gameRepository: gameRepository}
}

func (s *GameService) GetAllGames() ([]*entity.Game, error) {
	return s.gameRepository.GetAllGames()
}

func (s *GameService) AddGame(game restmodel.Game) (*restmodel.Game, error) {
	addedGame, err := s.gameRepository.AddGame(game)
	if err != nil {
		return nil, err
	}

	return addedGame, nil
}

func (s *GameService) UpdateGame(game entity.Game) (*entity.Game, error) {
	updatedGame, err := s.gameRepository.UpdateGame(game)
	if err != nil {
		return nil, err
	}
	return updatedGame, nil
}

func (s *GameService) ReleaseStockGame(order entity.Order) error {
	err := s.gameRepository.ReleaseStockGame(order)
	if err != nil {
		return fmt.Errorf("failed to Release Stock game ID %s: %w", order.ProductId, err)
	}

	return err
}

func (s *GameService) DeleteGame(id string) error {
	err := s.gameRepository.DeleteGame(id)
	if err != nil {
		return fmt.Errorf("failed to delete game with ID %s: %w", id, err)
	}

	return err
}
