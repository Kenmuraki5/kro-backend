package interfaces

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type GameService interface {
	GetAllGames() ([]*entity.Game, error)
	AddGame(restmodel.Game) (*restmodel.Game, error)
	UpdateGame(entity.Game) (*entity.Game, error)
	ReleaseStockGame(entity.Order) error
	DeleteGame(string) error
}
