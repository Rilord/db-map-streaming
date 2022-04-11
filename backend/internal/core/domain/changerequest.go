package domain

import (
	"database/sql"
)

type ChangeRequest struct {
	ID            string   `json:"id"`
	PinsID        []string `json:"pinids"`
	ItinerariesID []string
	CreatedAt     sql.NullTime
}
