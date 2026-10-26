package services

import (
	"context"
	"fmt"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/domain/models"
)

type gameService struct {
}

func NewGameService() interfaces.GameService {
	return &gameService{}
}

func (gs *gameService) CalculateNextMove(ctx context.Context, g *models.CurrentGame) models.Move {
	_, move := minimax(g.Game.Board, models.Player2, 0)
	return move
}

func (gs *gameService) Validation(ctx context.Context, g *models.CurrentGame, move models.Move) (bool, error) {
	if move.Row < 0 || move.Col < 0 || move.Row >= 3 || move.Col >= 3 {
		return false, fmt.Errorf("координаты должны быть от 0 до 2")
	}
	if g.Game.Board[move.Row][move.Col] != models.Empty {
		return false, fmt.Errorf("клетка [%d,%d] уже занята", move.Row, move.Col)
	}
	return true, nil
}

func (gs *gameService) CompletionCheck(ctx context.Context, g models.CurrentGame) models.GameStatus {
	return checkPlayResult(g.Game.Board, g.Status)
}

func checkPlayResult(b [3][3]int, res models.GameStatus) models.GameStatus {

	for j := 0; !isFinished(res) && j < 3; j++ {
		res = checkLine(b[j][0], b[j][1], b[j][2], res)
		res = checkLine(b[0][j], b[1][j], b[2][j], res)
	}

	res = checkLine(b[0][0], b[1][1], b[2][2], res)
	res = checkLine(b[2][0], b[1][1], b[0][2], res)

	if !isFinished(res) && checkDraw(b) {
		res = models.Draw
	}

	return res
}

func checkWinner(player int) models.GameStatus {
	if player == models.Player1 {
		return models.XWin
	}
	return models.OWin
}

func checkLine(a, b, c int, res models.GameStatus) models.GameStatus {
	if isFinished(res) {
		return res
	}
	if a != models.Empty && a == b && b == c {
		res = checkWinner(a)
	}
	return res
}

func checkDraw(board [3][3]int) bool {
	for _, v := range board {
		for _, b := range v {
			if b == models.Empty {
				return false
			}
		}
	}
	return true
}

func minimax(board [3][3]int, isPlayer int, depth int) (int, models.Move) {
	status := checkPlayResult(board, models.OTurn)
	if isFinished(status) {
		newStatus := checkStatus(status)
		if newStatus > 0 {
			newStatus = newStatus - depth
		} else if newStatus < 0 {
			newStatus = newStatus + depth
		}
		return newStatus, models.Move{Row: -1, Col: -1}
	}

	var bestScore int
	isBetter := func(current, best int) bool { return false }

	if isPlayer == models.Player1 {
		bestScore = -1000
		isBetter = func(s, b int) bool { return s > b }
	} else {
		bestScore = 1000
		isBetter = func(s, b int) bool { return s < b }
	}
	bestMove := models.Move{Row: -1, Col: -1}
	nextPlayer := models.Player1
	if isPlayer == models.Player1 {
		nextPlayer = models.Player2
	}

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r][c] == models.Empty {
				board[r][c] = isPlayer
				score, _ := minimax(board, nextPlayer, depth+1)
				board[r][c] = models.Empty

				if isBetter(score, bestScore) {
					bestScore = score
					bestMove.Row, bestMove.Col = r, c
				}
			}
		}
	}
	return bestScore, bestMove
}

func checkStatus(status models.GameStatus) int {
	if status == models.XWin {
		return 10
	}
	if status == models.OWin {
		return -10
	}
	return 0
}

func isFinished(s models.GameStatus) bool {
	return s == models.XWin || s == models.OWin || s == models.Draw
}
