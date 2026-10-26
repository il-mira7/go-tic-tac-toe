package repository

import (
	"context"
	"errors"
	"fmt"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/domain/models"
	"tic-tac-toe/internal/infrastructure/mappers"
	dto "tic-tac-toe/internal/infrastructure/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type gameRepository struct {
	db *pgxpool.Pool
}

func NewPostgresGameRepository(db *pgxpool.Pool) interfaces.GameRepository {
	return &gameRepository{db: db}
}

func (r *gameRepository) Save(ctx context.Context, game *models.CurrentGame) error {
	if game.CreatedAt.IsZero() {
		game.CreatedAt = time.Now()
	}

	dto, err := mappers.ToPostgres(game)
	if err != nil {
		return err
	}

	query := `
        INSERT INTO games (id, board, status, player_x_id, player_o_id, current_turn_id, winner_id, is_vs_computer, created_at, finished_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        ON CONFLICT (id) DO UPDATE SET 
            board = EXCLUDED.board, 
            status = EXCLUDED.status,
            player_o_id = EXCLUDED.player_o_id,
            current_turn_id = EXCLUDED.current_turn_id,
            winner_id = EXCLUDED.winner_id,
            finished_at = EXCLUDED.finished_at;
    `

	_, err = r.db.Exec(ctx, query,
		dto.ID,
		dto.Board,
		dto.Status,
		dto.PlayerXID,
		dto.PlayerOID,
		dto.CurrentTurnID,
		dto.WinnerID,
		dto.IsVsComputer,
		dto.CreatedAt,
		dto.FinishedAt,
	)

	return err
}

func (r *gameRepository) Get(ctx context.Context, id string) (*models.CurrentGame, error) {
	query := `SELECT id, board, status, player_x_id, player_o_id, current_turn_id, winner_id, is_vs_computer, created_at, finished_at
	          FROM games 
			  WHERE id = $1`

	gameUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}
	rows, err := r.db.Query(ctx, query, gameUUID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	gameDTO, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dto.GameDTO])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("game not found")
		}
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	return mappers.ToDomain(&gameDTO)
}

func (r *gameRepository) GetAvailableGames(ctx context.Context) ([]*models.CurrentGame, error) {
	query := `
			SELECT * FROM games 
			WHERE status = $1 AND player_o_id IS NULL ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query, "WAITING")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	dtos, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.GameDTO])
	if err != nil {
		return nil, fmt.Errorf("collect rows failed: %w", err)
	}

	games := make([]*models.CurrentGame, 0, len(dtos))
	for _, d := range dtos {
		domainGame, err := mappers.ToDomain(&d)
		if err != nil {
			return nil, fmt.Errorf("mapping failed: %w", err)
		}
		games = append(games, domainGame)
	}

	return games, nil
}

func (r *gameRepository) GetFinishedGamesByUserID(ctx context.Context, id string) ([]*models.CurrentGame, error) {
	query := `
			SELECT * FROM games 
			WHERE (player_x_id = $1 OR player_o_id = $1)
					AND (status = 'DRAW' OR status = 'X_WIN' OR status = 'O_WIN')
			ORDER BY created_at DESC;
			`
	gameUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}
	rows, err := r.db.Query(ctx, query, gameUUID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	dtos, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.GameDTO])
	if err != nil {
		return nil, fmt.Errorf("collect rows failed: %w", err)
	}

	games := make([]*models.CurrentGame, 0, len(dtos))
	for _, d := range dtos {
		domainGame, err := mappers.ToDomain(&d)
		if err != nil {
			return nil, fmt.Errorf("mapping failed: %w", err)
		}
		games = append(games, domainGame)
	}

	return games, nil
}

func (r *gameRepository) GetTopPlayers(ctx context.Context, limit int) ([]*models.Leaderboard, error) {
	query := `SELECT u.id as user_id, u.login,
					COUNT(CASE WHEN g.winner_id = u.id THEN 1 END)::FLOAT / 
					NULLIF(COUNT(g.id), 0) as win_ratio
				FROM users u
				JOIN games g ON (u.id = g.player_x_id OR u.id = g.player_o_id)
				WHERE g.status IN ('X_WIN', 'O_WIN', 'DRAW') AND g.is_vs_computer = FALSE  
				GROUP BY u.id, u.login
				ORDER BY win_ratio DESC
				LIMIT $1;
`

	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	dtos, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.LeaderboardDTO])
	if err != nil {
		return nil, fmt.Errorf("collect rows failed: %w", err)
	}

	leaders := make([]*models.Leaderboard, 0, len(dtos))
	for _, d := range dtos {
		leader := mappers.LeadersToDomain(&d)
		leaders = append(leaders, leader)
	}

	return leaders, nil
}
