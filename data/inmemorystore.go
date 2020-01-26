package data

import (
	"time"

	"github.com/dgravesa/drinklogs-service/model"
)

// InMemoryStore is a drink log store contained in memory.
type InMemoryStore struct {
	drinklogs map[uint64][]model.DrinkLog
}

// NewInMemoryStore returns a new in-memory drink log store.
func NewInMemoryStore() *InMemoryStore {
	var store InMemoryStore
	store.drinklogs = make(map[uint64][]model.DrinkLog)
	return &store
}

// Insert creates a new drink log for a user.
func (s *InMemoryStore) Insert(uid uint64, log model.DrinkLog) error {
	if _, exists := s.drinklogs[uid]; exists {
		s.drinklogs[uid] = append(s.drinklogs[uid], log)
	} else {
		s.drinklogs[uid] = []model.DrinkLog{log}
	}

	return nil
}

// InRange returns drink logs within a specified time range for a user.
func (s *InMemoryStore) InRange(uid uint64, ti, tf time.Time) []model.DrinkLog {
	result := []model.DrinkLog{}

	if userlogs, exists := s.drinklogs[uid]; exists {
		// add logs within time range to result
		for _, log := range userlogs {
			if log.Time.Before(tf) && log.Time.After(ti) {
				result = append(result, log)
			}
		}
	}

	return result
}
