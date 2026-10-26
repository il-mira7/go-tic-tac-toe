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

type userRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) interfaces.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, user *models.User) error {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	dto, err := mappers.UserToPostgres(user)
	if err != nil {
		return err
	}

	query := `
        INSERT INTO users (id, login, password_hash, created_at)
        VALUES ($1, $2, $3, $4);
    `
	_, err = r.db.Exec(ctx, query,
		dto.ID,
		dto.Login,
		dto.PasswordHash,
		dto.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("could not save user: %w", err)
	}

	return err
}

func (r *userRepository) Get(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, login, password_hash, created_at 
	          FROM users 
			  WHERE id = $1`

	gameUUID, _ := uuid.Parse(id)
	rows, err := r.db.Query(ctx, query, gameUUID)

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	userDTO, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dto.UserDTO])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	return mappers.UserToDomain(&userDTO), nil
}

func (r *userRepository) GetByLogin(ctx context.Context, login string) (*models.User, error) {
	query := `SELECT id, login, password_hash, created_at 
	          FROM users 
			  WHERE login = $1`

	rows, err := r.db.Query(ctx, query, login)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	userDTO, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dto.UserDTO])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	return mappers.UserToDomain(&userDTO), nil
}
