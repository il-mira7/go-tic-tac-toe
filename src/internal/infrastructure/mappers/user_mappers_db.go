package mappers

import (
	"tic-tac-toe/internal/domain/models"
	dto "tic-tac-toe/internal/infrastructure/models"

	"github.com/google/uuid"
)

func UserToPostgres(u *models.User) (*dto.UserDTO, error) {

	userID, err := uuid.Parse(u.ID)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID:           userID,
		Login:        u.Login,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
	}, nil
}

func UserToDomain(u *dto.UserDTO) *models.User {

	return &models.User{
		ID:           u.ID.String(),
		Login:        u.Login,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
	}
}
