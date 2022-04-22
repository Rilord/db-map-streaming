package interactors

import "game-server/internal/core/domain"

type ArticleInteractor interface {
	Add(article domain.Article) error
	Get(article domain.Article) (*domain.Article, error)
	Update(article domain.Article) error
	Pins(article_id string) ([]domain.Pin, error)
	Itineraries(article_id string) ([]domain.Itinerary, error)
	Delete(article_id string) error
}
