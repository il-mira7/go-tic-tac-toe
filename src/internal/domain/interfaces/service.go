package interfaces

import (
	"context"
	"tic-tac-toe/internal/domain/models"

	"github.com/google/uuid"
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
	Login(ctx context.Context, login, password string) (models.TokenPair, error)

	RefreshAccessToken(ctx context.Context, refreshToken string) (models.TokenPair, error)
	RefreshRefreshToken(ctx context.Context, refreshToken string) (models.TokenPair, error)
}

type TokenProvider interface {
	GenerateAccessToken(userID uuid.UUID) (string, error)
	GenerateRefreshToken(userID uuid.UUID) (string, error)
	ValidateAccessToken(tokenStr string) (bool, error)
	ValidateRefreshToken(tokenStr string) (bool, error)
	ExtractUUID(tokenStr string) (uuid.UUID, error)
	ExtractUUIDForRefresh(tokenStr string) (uuid.UUID, error)
}
