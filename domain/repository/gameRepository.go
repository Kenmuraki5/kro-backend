package repository

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type GameRepository interface {
	GetAllGames() ([]*entity.Game, error)
	AddGame(game restmodel.Game) (*restmodel.Game, error)
}
