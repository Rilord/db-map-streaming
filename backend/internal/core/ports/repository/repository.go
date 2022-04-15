package ports

import "game-server/internal/core/domain"

type MapTileRepository interface {
	Get(domain.Tile) (domain.Tile, error)
	GetRaster(tileId string) (*domain.Raster, error)
}

type VectorObjectRepository interface {
	Add(vector domain.VectorObject) error
	Get(domain.VectorObject) (domain.VectorObject, error)
	Delete(domain.VectorObject) error
}

type UserRepository interface {
	Add(domain.User) error
	Get(domain.User) (domain.User, error)
	Update(domain.User) error
	Delete(domain.User) error
}

type PinRepository interface {
	Add(domain.Pin) error
	Get(domain.Pin) (domain.Pin, error)
	Update(domain.Pin) error
	Delete(domain.Pin) error
}

type LocationRepository interface {
	Add(domain.Location) error
	Get(domain.Location) (domain.Location, error)
	Update(domain.Location) error
}

type ChangeRequestRepository interface {
	Add(domain.ChangeRequest) error
	Get(domain.ChangeRequest) (domain.ChangeRequest, error)
	Delete(domain.ChangeRequest) error
}

type ArticleRepository interface {
	Add(domain.Article) error
	Get(domain.Article) (domain.Article, error)
	GetPins(domain.Article) ([]domain.Pin, error)
	GetItineraries(domain.Article) ([]domain.Itinerary, error)
	GetComments(domain.Article) ([]domain.Comment, error)
	Delete(domain.Article) error
}

type ItineraryRepository interface {
	Add(itinerary *domain.Itinerary) error
	Get(itinerary_id string) (domain.Itinerary, error)
	Update(itinerary *domain.Itinerary) error
	Delete(itinerary_id string) error
}

type RoleRepository interface {
	FindUsersRole(domain.User) (domain.Role, error)
	FIndUsersWithRole(domain.Role) ([]domain.User, error)
}

type LayerRepository interface {
	Get(domain.Layer) (domain.Layer, error)
	GetTilesAtScale(domain.VectorObject) ([]domain.Tile, error)
}

type RasterRepository interface {
	Get(domain.Raster) (domain.Raster, error)
}

type CommentRepository interface {
	Add(domain.Comment) error
	Get(domain.Comment) (domain.Comment, error)
	Delete(domain.Comment) error
}
