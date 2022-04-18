package article

import "game-server/internal/core/domain"

func (i *Interactor) Add(article domain.Article) {
	i.Article.Add(article)
}
