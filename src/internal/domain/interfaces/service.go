package interfaces

import (
	"tic-tac-toe/internal/domain/models"
)

type GameService interface {
	CalculateNextMove(g *models.CurrentGame) models.Move
	Validation(g *models.CurrentGame, move models.Move) (bool, error)
	CompletionCheck(g models.CurrentGame) models.GameStatus
}
