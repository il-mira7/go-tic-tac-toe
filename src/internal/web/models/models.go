package models

type GameResponse struct {
	GameID string    `json:"game_id"`
	Board  [3][3]int `json:"board"`
	Status string    `json:"status"`
}

type GameUpdateRequest struct {
	Board [3][3]int `json:"board"`
	Row   int       `json:"row"`
	Col   int       `json:"col"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
