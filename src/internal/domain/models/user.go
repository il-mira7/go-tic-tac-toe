package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string
	Login        string
	PasswordHash string
	CreatedAt    time.Time
}

type SignUpRequest struct {
	Login    string
	Password string
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
	TokenType    string //  "Bearer"
}

type UserClaims struct {
	UserID uuid.UUID
}
