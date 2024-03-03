package repository

import (
	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
)

type ConsoleRepository interface {
	GetAllConsoles() ([]*entity.Console, error)
	AddConsole(restmodel.Console) (*restmodel.Console, error)
	UpdateConsole(entity.Console) (*entity.Console, error)
	UpdateStockConsole([]restmodel.Order) error
	ReleaseStockConsole(entity.Order) error
	DeleteConsole(string) error
}
