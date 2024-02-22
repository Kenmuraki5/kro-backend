package interfaces

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
)

type GameService interface {
	GetAllGames() ([]*entity.Game, error)
}
