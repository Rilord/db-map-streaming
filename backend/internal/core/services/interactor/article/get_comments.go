package article

import "game-server/internal/core/domain"

func (i *Interactor) GetComments(article_id string) ([]*domain.Comment, error) {
	return i.Article.GetComments(article_id)
}
