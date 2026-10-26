package services

import (
	"context"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/domain/models"
)

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx context.Context, user *models.User) error {
	return s.repo.Save(ctx, user)
}

func (s *userService) GetByLogin(ctx context.Context, login string) (*models.User, error) {
	return s.repo.GetByLogin(ctx, login)
}

func (s *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.Get(ctx, id)
}
