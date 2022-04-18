package interactors

import "game-server/internal/core/domain"

type ItineraryInteratror interface {
	Add(itinerary domain.Itinerary) error
	Get(itinerary domain.Itinerary) (domain.Itinerary, error)
	Update(itinerary domain.Itinerary) error
	Delete(itinerary_id string) error
}
