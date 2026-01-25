package mappers

import (
	"tic-tac-toe/internal/domain/models"
	"tic-tac-toe/internal/infrastructure/storage"
)

func ToStorageGame(g models.CurrentGame) storage.GameRecord {
	return storage.GameRecord{
		ID:    g.ID,
		Board: g.Game.Board}
}
func ToDomainGame(g storage.GameRecord) models.CurrentGame {

	return models.CurrentGame{
		ID: g.ID,
		Game: models.GameBoard{
			Board: g.Board}}
}
