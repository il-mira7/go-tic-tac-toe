package mappers

import (
	domain "tic-tac-toe/internal/domain/models"
	web "tic-tac-toe/internal/web/models"
)

func ToDomain(id string, req web.GameUpdateRequest) domain.CurrentGame {
	return domain.CurrentGame{
		ID: id,
		Game: domain.GameBoard{
			Board: req.Board,
		},
	}
}

func ToResponse(dg domain.CurrentGame, status domain.GameStatus) web.GameResponse {
	var stringStatus string
	switch status {
	case domain.Ongoing:
		stringStatus = "Ongoing"
	case domain.Draw:
		stringStatus = "Draw"
	case domain.Player1Wins:
		stringStatus = "Player 1 Wins"
	case domain.Player2Wins:
		stringStatus = "Computer Wins"
	default:
		stringStatus = "Unknown"
	}
	return web.GameResponse{
		GameID: dg.ID,
		Board:  dg.Game.Board,
		Status: stringStatus,
	}
}
