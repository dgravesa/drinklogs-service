package data

import "time"

import "github.com/dgravesa/drinklogs-service/model"

// DrinkLogStore is a store interface for user drink log data.
type DrinkLogStore interface {
	Insert(uid uint64, log model.DrinkLog) error
	InRange(uid uint64, ti, tf time.Time) []model.DrinkLog
}

var drinkLogStore DrinkLogStore

// SetDrinkLogStore sets the drink log data backend.
func SetDrinkLogStore(store DrinkLogStore) {
	drinkLogStore = store
}

// DrinkLogsInRange returns all drink logs in a time range for a user.
func DrinkLogsInRange(uid uint64, ti, tf time.Time) []model.DrinkLog {
	return drinkLogStore.InRange(uid, ti, tf)
}

// InsertDrinkLog inserts a new drink log for a user.
func InsertDrinkLog(uid uint64, log model.DrinkLog) error {
	return drinkLogStore.Insert(uid, log)
}
