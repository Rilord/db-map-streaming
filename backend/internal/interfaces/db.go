package interfaces

import "game-server/model"

type MapTilePaginationConfig struct {
	First *int
	After *string
}

type MapTileDB interface {
	Get(id string) (model.Tile, error)
	Pagination(config MapTilePaginationConfig) error
}

type VectorObjectDB interface {
	Get(id string) (model.VectorObject, error)
	Create(object *model.VectorObject) error
}

type UserPaginationConfig struct {
	model.UserFilter
	First *int
	After *string
}

type UserDB interface {
	Get(id string) (model.User, error)
	Create(user *model.User) error
	Pagination(config UserPaginationConfig) (model.UserConnection, error)
}

type PinDB interface {
	Get(id string) (mode.Pin, error)
	Create(pin *model.Pin) error
	Update(pin *model.Pin) error
}

type ChangeRequestDB interface {
	Get(id string) (model.ChangeRequest, error)
	Create(request *model.ChangeRequest) error
	Merge(id string) error
}

type ItineraryDB interface {
	Get(id string) (model.Itinerary, error)
	Create(itinerary *model.Itinerary) error
	Update(itinerary *model.Itinerary) error
}

type RoleFindConfig struct {
	UserIDs []string
	Types   []model.RoleType
	Limit   int
}

type RoleDB interface {
	Find(config RoleFindConfig) ([]model.Role, error)
}

type LayerDB interface {
	Get(id string) (model.Layer, error)
}

type RasterDB interface {
	Get(id string) (model.Raster, error)
}

type db interface {
	User() UserDB
	Role() RoleDB
	Pin() PinDB
	Itinerary() ItineraryDB
}
