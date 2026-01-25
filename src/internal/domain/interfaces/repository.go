package interfaces

import (
	"tic-tac-toe/internal/domain/models"
)

type GameRepository interface {
	Save(game models.CurrentGame) error
	Get(id string) (models.CurrentGame, error)
}
