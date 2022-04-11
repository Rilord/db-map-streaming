package ports

import "game-server/internal/core/domain"

type MapTileDB interface {
	Get(id string) (domain.Tile, error)
}

type VectorObjectDB interface {
	Get(id string) (domain.VectorObject, error)
	Create(object *domain.VectorObject) error
}

type UserDB interface {
	Get(id string) (domain.User, error)
	Create(user *domain.User) error
}

type PinDB interface {
	Get(id string) (domain.Pin, error)
	Create(pin *domain.Pin) error
	Update(pin *domain.Pin) error
}

type LocationDB interface {
	Get(id string) (domain.Location, error)
	Create(location *domain.Location) error
	Update(location *domain.Location) error
}

type ChangeRequestDB interface {
	Get(id string) (domain.ChangeRequest, error)
	Create(request *domain.ChangeRequest) error
	Delete(id string) error
}

type ArticleDB interface {
	Get(id string) (domain.Article, error)
	Create(article *domain.Article) error
	Delete(id string) error
}

type ItineraryDB interface {
	Get(id string) (domain.Itinerary, error)
	Create(itinerary *domain.Itinerary) error
	Update(itinerary *domain.Itinerary) error
}

type RoleFindConfig struct {
	UserIDs []string
	Types   []domain.RoleType
	Limit   int
}

type RoleDB interface {
	Find(config RoleFindConfig) ([]domain.Role, error)
}

type LayerDB interface {
	Get(id string) (domain.Layer, error)
}

type RasterDB interface {
	Get(id string) (domain.Raster, error)
}

type DB interface {
	User() UserDB
	Role() RoleDB
	Pin() PinDB
	ChangeRequest() ChangeRequestDB
	Layer() LayerDB
	Raster() RasterDB
	Location() LocationDB
	Itinerary() ItineraryDB
}
