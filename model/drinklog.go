package model

import "time"

// DrinkLog contains a time/amount pair representing a drink event.
type DrinkLog struct {
	Amount float64   `json:"amount"`
	Time   time.Time `json:"time"`
}
