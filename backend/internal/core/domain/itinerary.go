package domain

type ItineraryType string

const (
	UsersItinerary   ItineraryType = "USER_ITINERARY"
	DefaultItinerary ItineraryType = "DEFAULT_ITINERARY"
)

type Itinerary struct {
	ID       string `json:"id"`
	name     string
	Type     ItineraryType
	tags     string
	authorID string
	fromID   string
	toID     string
}
