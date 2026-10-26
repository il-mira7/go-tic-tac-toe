package mappers

import (
	"tic-tac-toe/internal/domain/models"
	webModels "tic-tac-toe/internal/web/models"
)

func ToDomainSignUp(req webModels.SignUpRequest) models.SignUpRequest {
	return models.SignUpRequest{
		Login:    req.Login,
		Password: req.Password,
	}
}

func ToWebUserInfoResponse(u *models.User) webModels.UserInfoResponse {
	return webModels.UserInfoResponse{
		ID:        u.ID,
		Login:     u.Login,
		CreatedAt: u.CreatedAt.String(),
	}
}
