package ports

import "game-server/internal/core/domain"

type MapTileFilter struct {
	Ids   []string
	Limit int
}

type MapTileRepository interface {
	Get(tile_id string) (domain.Tile, error)
	GetRaster(tile_id string) (*domain.Raster, error)
	GetVector(tile_id string) ([]*domain.Raster, error)
	Find(filter MapTileFilter) ([]*domain.Tile, error)
}

type VectorObjectFilter struct {
	Ids   []string
	Type  int
	Limit int
}

type VectorObjectRepository interface {
	Add(vector domain.VectorObject) error
	Get(vector_id string) (domain.VectorObject, error)
	Find(filter VectorObjectFilter) ([]*domain.VectorObject, error)
	Delete(vector_id string) error
}

type UserRepository interface {
	Add(domain.User) error
	Get(user_id string) (domain.User, error)
	Update(domain.User) error
	Delete(user_id string) error
}

type PinFilter struct {
	Ids   []string
	Type  string
	Limit int
	tags  []string
}

type PinRepository interface {
	Add(domain.Pin) error
	Get(pin_id string) (domain.Pin, error)
	Find(filter PinFilter) ([]*domain.Pin, error)
	Update(domain.Pin) error
	Delete(pin_id string) error
}

type LocationRepository interface {
	Add(domain.Location) error
	Get(location_id string) (domain.Location, error)
	Update(domain.Location) error
}

type ChangeRequestRepository interface {
	Add(domain.ChangeRequest) error
	Get(request_id string) (domain.ChangeRequest, error)
	Delete(domain.ChangeRequest) error
}

type ArticleRepository interface {
	Add(domain.Article) error
	Get(article_id string) (domain.Article, error)
	GetPins(article_id string) ([]*domain.Pin, error)
	GetItineraries(article_id string) ([]*domain.Itinerary, error)
	GetChangeRequests(article_id string) ([]*domain.ChangeRequest, error)
	GetComments(article_id string) ([]*domain.Comment, error)
	Delete(article_id string) error
}

type ItineraryRepository interface {
	Add(itinerary *domain.Itinerary) error
	Get(itinerary_id string) (domain.Itinerary, error)
	Update(itinerary *domain.Itinerary) error
	Delete(itinerary_id string) error
}

type RoleFilter struct {
	Ids   []string
	Users []string
	Types []domain.RoleType
}

type RoleRepository interface {
	Find(filter RoleFilter) ([]*domain.User, error)
}

type LayerRepository interface {
	Get(layer_id string) (domain.Layer, error)
}

type RasterRepository interface {
	Get(raster_id string) (domain.Raster, error)
}

type CommentFilter struct {
	Ids       []string
	ArticleId string
	CreatorId string
}

type CommentRepository interface {
	Add(domain.Comment) error
	Get(comment_id string) (domain.Comment, error)
	Find(filter CommentFilter) ([]*domain.Comment, error)
	Delete(comment_id string) error
}
