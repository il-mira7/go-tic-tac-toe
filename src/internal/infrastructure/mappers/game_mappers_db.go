package mappers

import (
	"encoding/json"
	"tic-tac-toe/internal/domain/models"
	dto "tic-tac-toe/internal/infrastructure/models"
	"time"

	"github.com/google/uuid"
)

func ToPostgres(g *models.CurrentGame) (*dto.GameDTO, error) {

	board, err := json.Marshal(g.Game.Board)
	if err != nil {
		return nil, err
	}

	gameID, err := uuid.Parse(g.ID)
	if err != nil {
		return nil, err
	}

	playerX, err := uuid.Parse(g.PlayerXID)
	if err != nil {
		return nil, err
	}

	return &dto.GameDTO{
		ID:            gameID,
		Board:         board,
		Status:        statusToString(g.Status),
		PlayerXID:     playerX,
		PlayerOID:     parseUUIDPtr(g.PlayerOID),
		CurrentTurnID: parseUUIDPtr(g.CurrentTurnID),
		WinnerID:      parseUUIDPtr(g.WinnerID),
		IsVsComputer:  g.IsVsComputer,
		CreatedAt:     g.CreatedAt,
		FinishedAt:    timeToPtr(g.FinishedAt),
	}, nil
}

func ToDomain(g *dto.GameDTO) (*models.CurrentGame, error) {
	var board [3][3]int

	if err := json.Unmarshal(g.Board, &board); err != nil {
		return nil, err
	}

	return &models.CurrentGame{
		ID: g.ID.String(),
		Game: models.GameBoard{
			Board: board,
		},
		Status:        stringStatusToGameStatus(g.Status),
		PlayerXID:     g.PlayerXID.String(),
		PlayerOID:     uuidPtrToString(g.PlayerOID),
		CurrentTurnID: uuidPtrToString(g.CurrentTurnID),
		WinnerID:      uuidPtrToString(g.WinnerID),
		IsVsComputer:  g.IsVsComputer,
		CreatedAt:     g.CreatedAt,
		FinishedAt:    timePtrToValue(g.FinishedAt),
	}, nil
}

func parseUUIDPtr(s string) *uuid.UUID {
	id, err := uuid.Parse(s)
	if err != nil {
		return nil
	}
	return &id
}

func uuidPtrToString(u *uuid.UUID) string {
	if u == nil {
		return ""
	}
	return u.String()
}

func statusToString(s models.GameStatus) string {
	switch s {
	case models.Waiting:
		return "WAITING"
	case models.XTurn:
		return "X_TURN"
	case models.OTurn:
		return "O_TURN"
	case models.XWin:
		return "X_WIN"
	case models.OWin:
		return "O_WIN"
	case models.Draw:
		return "DRAW"
	default:
		return "WAITING"
	}
}

func stringStatusToGameStatus(s string) models.GameStatus {
	switch s {
	case "WAITING":
		return models.Waiting
	case "X_TURN":
		return models.XTurn
	case "O_TURN":
		return models.OTurn
	case "X_WIN":
		return models.XWin
	case "O_WIN":
		return models.OWin
	case "DRAW":
		return models.Draw
	default:
		return models.Waiting
	}
}

func timePtrToValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func timeToPtr(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}
