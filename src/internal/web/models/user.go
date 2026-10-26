package models

type SignUpRequest struct {
	Login    string `json:"login" binding:"required,min=3"`    //  логин минимум 3 символа
	Password string `json:"password" binding:"required,min=6"` // Пароль минимум 6 символов
}

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
