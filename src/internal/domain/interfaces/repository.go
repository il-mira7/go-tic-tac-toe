package interfaces

import (
	"context"
	"tic-tac-toe/internal/domain/models"
)

type GameRepository interface {
	Save(ctx context.Context, game *models.CurrentGame) error
	Get(ctx context.Context, id string) (*models.CurrentGame, error)
	GetAvailableGames(ctx context.Context) ([]*models.CurrentGame, error)
}

type UserRepository interface {
	Save(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id string) (*models.User, error)
	GetByLogin(ctx context.Context, login string) (*models.User, error)
}
