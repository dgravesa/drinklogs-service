package data

import "time"

import "github.com/dgravesa/drinklogs-service/model"

// DrinkLogStore is a store interface for user drink log data.
type DrinkLogStore interface {
	// Insert creates a new drink log for a user.
	Insert(uid string, log model.DrinkLog) error

	// InRange returns drink logs within a specified time range for a user.
	InRange(uid string, ti, tf time.Time) ([]model.DrinkLog, error)
}

var drinkLogStore DrinkLogStore

// SetDrinkLogStore sets the drink log data backend.
func SetDrinkLogStore(store DrinkLogStore) {
	drinkLogStore = store
}

// DrinkLogsInRange returns all drink logs in a time range for a user.
func DrinkLogsInRange(uid string, ti, tf time.Time) ([]model.DrinkLog, error) {
	return drinkLogStore.InRange(uid, ti, tf)
}

// InsertDrinkLog inserts a new drink log for a user.
func InsertDrinkLog(uid string, log model.DrinkLog) error {
	return drinkLogStore.Insert(uid, log)
}
