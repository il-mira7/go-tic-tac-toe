package storage

import (
	"fmt"
)

func NewGameStorage() *GameStorage {
	return &GameStorage{
		games: make(map[string]GameRecord),
	}
}

func (gs *GameStorage) Save(id string, game GameRecord) error {
	if id == "" {
		return fmt.Errorf("game ID cannot be empty")
	}

	gs.mutex.Lock()
	gs.games[id] = game
	gs.mutex.Unlock()

	return nil
}

func (gs *GameStorage) Get(id string) (GameRecord, error) {
	if id == "" {
		return GameRecord{}, fmt.Errorf("game ID cannot be empty")
	}

	gs.mutex.RLock()
	game, exists := gs.games[id]
	gs.mutex.RUnlock()

	if !exists {
		return GameRecord{}, fmt.Errorf("game with ID %s not found", id)
	}

	return game, nil
}
