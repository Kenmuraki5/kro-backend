// services/game_service.go

package services

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/repository"
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
