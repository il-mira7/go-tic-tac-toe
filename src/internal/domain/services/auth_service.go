package services

import (
	"context"
	"errors"
	"fmt"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/domain/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userSvc     interfaces.UserService
	jwtProvider interfaces.TokenProvider
}

func NewAuthService(userSvc interfaces.UserService, jwtProv interfaces.TokenProvider) interfaces.AuthService {
	return &authService{
		userSvc:     userSvc,
		jwtProvider: jwtProv,
	}
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

func (s *authService) Login(ctx context.Context, login, password string) (models.TokenPair, error) {

	user, err := s.userSvc.GetByLogin(ctx, login)
	if err != nil {
		return models.TokenPair{}, fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return models.TokenPair{}, fmt.Errorf("invalid password")
	}

	userUUID, err := uuid.Parse(user.ID)
	if err != nil {
		return models.TokenPair{}, fmt.Errorf("invalid user id format in database")
	}

	access, err := s.jwtProvider.GenerateAccessToken(userUUID)
	if err != nil {
		return models.TokenPair{}, fmt.Errorf("failed to generate access token: %w", err)
	}

	refresh, err := s.jwtProvider.GenerateRefreshToken(userUUID)
	if err != nil {
		return models.TokenPair{}, fmt.Errorf("failed to generate access token: %w", err)
	}

	return models.TokenPair{
		AccessToken:  access,
		RefreshToken: refresh,
		TokenType:    "Bearer",
	}, nil
}

func (s *authService) RefreshAccessToken(ctx context.Context, refreshToken string) (models.TokenPair, error) {
	userID, err := s.jwtProvider.ExtractUUIDForRefresh(refreshToken)
	if err != nil {
		return models.TokenPair{}, errors.New("invalid refresh token")
	}

	newAccess, err := s.jwtProvider.GenerateAccessToken(userID)
	if err != nil {
		return models.TokenPair{}, err
	}

	return models.TokenPair{
		AccessToken:  newAccess,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}, nil
}

func (s *authService) RefreshRefreshToken(ctx context.Context, refreshToken string) (models.TokenPair, error) {
	userID, err := s.jwtProvider.ExtractUUIDForRefresh(refreshToken)
	if err != nil {
		return models.TokenPair{}, errors.New("invalid refresh token")
	}

	newAccess, _ := s.jwtProvider.GenerateAccessToken(userID)
	newRefresh, _ := s.jwtProvider.GenerateRefreshToken(userID)

	return models.TokenPair{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
		TokenType:    "Bearer",
	}, nil
}
