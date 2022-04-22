package article

import (
	"game-server/internal/core/domain"
)

func (i *Interactor) Update(article domain.Article) {
	i.Article.Update(article)
}
