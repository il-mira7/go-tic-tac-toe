package repository

import (
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/domain/models"
	"tic-tac-toe/internal/infrastructure/mappers"
	"tic-tac-toe/internal/infrastructure/storage"
)

type gameRepository struct {
	storage *storage.GameStorage
}

func NewGameRepository(storage *storage.GameStorage) interfaces.GameRepository {
	return &gameRepository{storage: storage}
}

func (r *gameRepository) Save(game models.CurrentGame) error {
	storageRecord := mappers.ToStorageGame(game)
	return r.storage.Save(game.ID, storageRecord)
}

func (r *gameRepository) Get(id string) (models.CurrentGame, error) {
	storageRecord, err := r.storage.Get(id)
	if err != nil {
		return models.CurrentGame{}, err
	}

	domainGame := mappers.ToDomainGame(storageRecord)

	return domainGame, nil
}
