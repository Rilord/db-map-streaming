package article

import (
	pi "game-server/internal/core/ports/interactors/article"
	sa "game-server/internal/core/services/article"
)

func New() pi.ArticleInteractor {
	return sa.article.New()
}
