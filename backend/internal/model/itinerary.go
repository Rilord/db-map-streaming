package model

import "gorm.io/gorm"

type ItineraryType string

const (
	UsersItinerary   ItineraryType = "USER_ITINERARY"
	DefaultItinerary ItineraryType = "DEFAULT_ITINERARY"
)

type Itinerary struct {
	gorm.Model
	ID       string `gorm:"default:generated();" json:"id"`
	name     string
	Type     ItineraryType
	tags     string
	authorID string
	fromID   string
	toID     string
}
