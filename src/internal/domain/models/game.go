package models

type GameBoard struct {
	Board [3][3]int
}

const (
	Empty   = 0
	Player1 = 1
	Player2 = 2
)

type CurrentGame struct {
	ID   string
	Game GameBoard
}

type GameStatus int

const (
	Ongoing GameStatus = iota
	Draw
	Player1Wins
	Player2Wins
)

type Move struct {
	Row, Col int
}
