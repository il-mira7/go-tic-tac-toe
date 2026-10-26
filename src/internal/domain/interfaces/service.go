package interfaces

import (
	"context"
	"tic-tac-toe/internal/domain/models"
)

type GameService interface {
	CalculateNextMove(ctx context.Context, g *models.CurrentGame) models.Move
	Validation(ctx context.Context, g *models.CurrentGame, move models.Move) (bool, error)
	CompletionCheck(ctx context.Context, g models.CurrentGame) models.GameStatus
}

type UserService interface {
	Create(ctx context.Context, user *models.User) error
	GetByLogin(ctx context.Context, login string) (*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
}

type AuthService interface {
	SignUp(ctx context.Context, req models.SignUpRequest) error
	Authenticate(ctx context.Context, login, password string) (string, error)
}
