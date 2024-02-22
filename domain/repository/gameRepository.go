package repository

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
)

type GameRepository interface {
	GetAllGames() ([]*entity.Game, error)
}
