package models

import "time"

type GameBoard struct {
	Board [3][3]int
}

const (
	Empty   = 0
	Player1 = 1
	Player2 = 2
)

type CurrentGame struct {
	ID     string
	Game   GameBoard
	Status GameStatus

	PlayerXID string
	PlayerOID string

	CurrentTurnID string
	WinnerID      string
	IsVsComputer  bool

	CreatedAt  time.Time
	FinishedAt time.Time
}

type GameStatus int

const (
	Waiting GameStatus = iota
	XTurn
	OTurn

	XWin
	OWin
	Draw
)

type Move struct {
	Row, Col int
}
