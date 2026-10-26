package models

type GameResponse struct {
	GameID        string    `json:"game_id"`
	Board         [3][3]int `json:"board"`
	Status        string    `json:"status"`
	PlayerXID     string    `json:"player_x_id"`
	PlayerOID     string    `json:"player_o_id"`
	CurrentTurnID string    `json:"current_turn_id"`
	WinnerID      string    `json:"winner_id,omitempty"`
}

type GameUpdateRequest struct {
	Board [3][3]int `json:"board"`
	Row   int       `json:"row"`
	Col   int       `json:"col"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type LeaderBoardResponse struct {
	ID       string
	Login    string
	WinRatio float64
}
