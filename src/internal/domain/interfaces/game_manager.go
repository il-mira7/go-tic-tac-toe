package interfaces

import (
	"context"
	"tic-tac-toe/internal/domain/models"
)

type GameManager interface {
	CreateGame(ctx context.Context, playerXID string, isVsBot bool) (*models.CurrentGame, error)
	JoinGame(ctx context.Context, gameID string, playerOID string) error

	MakeMove(ctx context.Context, gameID string, playerID string, move models.Move) (*models.CurrentGame, error)
	GetAvailableGames(ctx context.Context) ([]*models.CurrentGame, error)
	GetGame(ctx context.Context, id string) (*models.CurrentGame, error)
}

type GameInfoManager interface {
	GetUserHistory(ctx context.Context, id string) ([]*models.CurrentGame, error)
	GetLeaderboard(ctx context.Context, limit int) ([]*models.Leaderboard, error)
}
