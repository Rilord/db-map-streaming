package article

import "game-server/internal/core/domain"

func (i *Interactor) Add(article domain.Article) (string, error) {
	return i.Article.Add(article)
}
