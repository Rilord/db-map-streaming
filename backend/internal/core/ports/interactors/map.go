package interactors

import "game-server/internal/core/domain"

type MapTileInteractor interface {
	Get(tile domain.Tile) (domain.Tile, error)
	Pins(tile_id string) ([]domain.Tile, error)
	Itineraries(tile_id string) ([]domain.Itinerary, error)
	Raster(tile_id string) (domain.Raster, error)
}
