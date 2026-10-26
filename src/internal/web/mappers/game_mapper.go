package mappers

import (
	"tic-tac-toe/internal/domain/models"
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

func ToWebResponse(dg *domain.CurrentGame) web.GameResponse {
	return web.GameResponse{
		GameID:        dg.ID,
		Board:         dg.Game.Board,
		Status:        statusToWebString(dg.Status),
		PlayerXID:     dg.PlayerXID,
		PlayerOID:     dg.PlayerOID,
		CurrentTurnID: dg.CurrentTurnID,
		WinnerID:      dg.WinnerID,
	}
}

func ToWebResponseList(games []*models.CurrentGame) []web.GameResponse {
	res := make([]web.GameResponse, 0, len(games))
	for _, g := range games {
		res = append(res, ToWebResponse(g))
	}
	return res
}

func LeadersToWeb(users []*models.Leaderboard) []web.LeaderBoardResponse {
	toWeb := func(u *models.Leaderboard) web.LeaderBoardResponse {
		return web.LeaderBoardResponse{
			ID:       u.ID,
			Login:    u.Login,
			WinRatio: u.WinRatio,
		}
	}

	res := make([]web.LeaderBoardResponse, 0, len(users))
	for _, user := range users {
		res = append(res, toWeb(user))
	}
	return res
}

func statusToWebString(s domain.GameStatus) string {
	switch s {
	case domain.Waiting:
		return "Waiting for player"
	case domain.XTurn:
		return "Player X Turn"
	case domain.OTurn:
		return "Player O Turn"
	case domain.XWin:
		return "Player X Wins"
	case domain.OWin:
		return "Player O Wins"
	case domain.Draw:
		return "Draw"
	default:
		return "Unknown"
	}
}
