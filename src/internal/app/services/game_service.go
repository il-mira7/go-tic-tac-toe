package services

import (
	"context"
	"errors"
	"fmt"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type repositoryGameManager struct {
	repository    interfaces.GameRepository
	domainService interfaces.GameService
}

func NewRepositoryGameUseManager(repo interfaces.GameRepository, domainSvc interfaces.GameService) interfaces.GameManager {
	return &repositoryGameManager{
		repository:    repo,
		domainService: domainSvc,
	}
}

func (g *repositoryGameManager) CreateGame(ctx context.Context, playerXID string, isVsBot bool) (*models.CurrentGame, error) {
	game := &models.CurrentGame{
		ID:           uuid.NewString(),
		PlayerXID:    playerXID,
		Status:       models.Waiting,
		IsVsComputer: isVsBot,
		CreatedAt:    time.Now(),
	}

	if isVsBot {
		game.PlayerOID = uuid.Nil.String()
		game.Status = models.XTurn
		game.CurrentTurnID = game.PlayerXID
	}

	err := g.repository.Save(ctx, game)
	return game, err
}

func (g *repositoryGameManager) JoinGame(ctx context.Context, gameID string, playerOID string) error {
	game, err := g.repository.Get(ctx, gameID)
	if err != nil {
		return err
	}

	if game.Status != models.Waiting {
		return errors.New("game is already full")
	}

	game.PlayerOID = playerOID
	game.Status = models.XTurn
	game.CurrentTurnID = game.PlayerXID

	return g.repository.Save(ctx, game)
}

func (g *repositoryGameManager) MakeMove(ctx context.Context, gameID string, playerID string, move models.Move) (*models.CurrentGame, error) {
	game, err := g.repository.Get(ctx, gameID)
	if err != nil {
		return nil, err
	}

	if game.Status == models.XWin || game.Status == models.OWin || game.Status == models.Draw {
		return nil, errors.New("game is already finished")
	}
	if game.CurrentTurnID != playerID {
		return nil, errors.New("it is not your turn")
	}

	ok, err := g.domainService.Validation(ctx, game, move)
	if !ok {
		return nil, err
	}

	symbol := models.Player1
	if playerID == game.PlayerOID {
		symbol = models.Player2
	}
	game.Game.Board[move.Row][move.Col] = symbol

	game.Status = g.domainService.CompletionCheck(ctx, *game)

	if game.Status == models.XTurn || game.Status == models.OTurn {
		if game.IsVsComputer {
			botMove := g.domainService.CalculateNextMove(ctx, game)
			game.Game.Board[botMove.Row][botMove.Col] = models.Player2
			game.Status = g.domainService.CompletionCheck(ctx, *game)
			game.CurrentTurnID = game.PlayerXID
		} else {
			if playerID == game.PlayerXID {
				game.CurrentTurnID = game.PlayerOID
				game.Status = models.OTurn
			} else {
				game.CurrentTurnID = game.PlayerXID
				game.Status = models.XTurn
			}
			game.Status = g.domainService.CompletionCheck(ctx, *game)
		}
	}

	switch game.Status {
	case models.XWin:
		game.WinnerID = game.PlayerXID
		game.FinishedAt = time.Now()
	case models.OWin:
		game.WinnerID = game.PlayerOID
		game.FinishedAt = time.Now()
	case models.Draw:
		game.FinishedAt = time.Now()
	}

	err = g.repository.Save(ctx, game)
	if err != nil {
		return nil, err
	}

	return game, nil

}

func (g *repositoryGameManager) GetAvailableGames(ctx context.Context) ([]*models.CurrentGame, error) {
	games, err := g.repository.GetAvailableGames(ctx)
	if err != nil {
		return nil, err
	}

	if games == nil {
		return []*models.CurrentGame{}, nil
	}

	return games, nil
}

func (g *repositoryGameManager) GetGame(ctx context.Context, id string) (*models.CurrentGame, error) {
	game, err := g.repository.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find game in database: %w", err)
	}

	return game, nil
}
