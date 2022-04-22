package article

import (
	"game-server/internal/core/domain"
)

func (i *Interactor) Get(article_id string) (domain.Article, error) {
	return i.Article.Get(article_id)
}
