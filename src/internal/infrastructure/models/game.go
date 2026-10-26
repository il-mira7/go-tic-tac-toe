package dto

import (
	"time"

	"github.com/google/uuid"
)

type GameDTO struct {
	ID     uuid.UUID `db:"id"`
	Board  []byte    `db:"board"`
	Status string    `db:"status"`

	PlayerXID     uuid.UUID  `db:"player_x_id"`
	PlayerOID     *uuid.UUID `db:"player_o_id"`
	CurrentTurnID *uuid.UUID `db:"current_turn_id"`
	WinnerID      *uuid.UUID `db:"winner_id"`

	IsVsComputer bool `db:"is_vs_computer"`

	CreatedAt  time.Time  `db:"created_at"`
	FinishedAt *time.Time `db:"finished_at"`
}
