package article

import "game-server/internal/core/domain"

func (i *Interactor) GetChangeRequests(article_id string) ([]*domain.ChangeRequest, error) {
	return i.Article.GetChangeRequests(article_id)
}
