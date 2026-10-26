package database

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed init.sql
var initSQL string

func NewPostgresPool(connStr string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	if err := initSchema(ctx, pool); err != nil {
		return nil, fmt.Errorf("failed to init schema: %w", err)
	}

	return pool, nil
}

func initSchema(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, initSQL)
	if err != nil {
		return fmt.Errorf("execute init sql: %w", err)
	}
	return nil
}
