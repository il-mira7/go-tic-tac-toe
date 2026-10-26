package services

import (
	"context"
	"fmt"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/domain/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userSvc interfaces.UserService
}

func NewAuthService(userSvc interfaces.UserService) interfaces.AuthService {
	return &authService{userSvc: userSvc}
}

func (s *authService) SignUp(ctx context.Context, req models.SignUpRequest) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		ID:           uuid.NewString(),
		Login:        req.Login,
		PasswordHash: string(hashed),
	}

	return s.userSvc.Create(ctx, user)
}

func (s *authService) Authenticate(ctx context.Context, login, password string) (string, error) {
	user, err := s.userSvc.GetByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	return user.ID, nil
}
