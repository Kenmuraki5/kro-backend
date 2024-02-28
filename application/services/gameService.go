// services/game_service.go

package services

import (
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
