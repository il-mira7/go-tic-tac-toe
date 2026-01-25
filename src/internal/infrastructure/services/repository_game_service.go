package services

import (
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/domain/models"
)

type repositoryGameService struct {
	repository    interfaces.GameRepository
	domainService interfaces.GameService
}

func NewRepositoryGameService(repo interfaces.GameRepository, domainSvc interfaces.GameService) interfaces.GameService {
	return &repositoryGameService{
		repository:    repo,
		domainService: domainSvc,
	}
}

func (rgs *repositoryGameService) CalculateNextMove(g *models.CurrentGame) models.Move {
	move := rgs.domainService.CalculateNextMove(g)

	g.Game.Board[move.Row][move.Col] = models.Player2 // компьтер
	rgs.repository.Save(*g)

	return move
}

func (rgs *repositoryGameService) Validation(g *models.CurrentGame, move models.Move) (bool, error) {
	savedGame, err := rgs.repository.Get(g.ID)
	if err == nil {
		g.Game.Board = savedGame.Game.Board
	}

	ok, err := rgs.domainService.Validation(g, move)
	if ok {
		g.Game.Board[move.Row][move.Col] = models.Player1
		rgs.repository.Save(*g)
	}
	return ok, err
}

func (rgs *repositoryGameService) CompletionCheck(g models.CurrentGame) models.GameStatus {
	return rgs.domainService.CompletionCheck(g)
}
