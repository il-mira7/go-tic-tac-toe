package storage

import (
	"sync"
)

type GameRecord struct {
	ID    string
	Board [3][3]int
}

type GameStorage struct {
	games map[string]GameRecord
	mutex sync.RWMutex
}
