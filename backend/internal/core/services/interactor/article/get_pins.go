package article

import "game-server/internal/core/domain"

func (i *Interactor) GetPins(article_id string) ([]*domain.Pin, error) {
	return i.Article.GetPins(article_id)
}
