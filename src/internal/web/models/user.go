package models

type RegisterResponse struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	UUID string `json:"uuid"`
}

type UserInfoResponse struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	CreatedAt string `json:"created_at"`
}

type JwtRequest struct {
	Login    string `json:"login" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3"`
}

type JwtResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
}

type RefreshJwtRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
