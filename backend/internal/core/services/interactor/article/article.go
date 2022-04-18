package article

import (
	ar "game-server/internal/adapter/repository"
	pr "game-server/internal/core/ports/repository"
)

type Interactor struct {
	Article pr.ArticleRepository
}

func NewInteractor() *Interactor {
	interactor := Interactor{
		Article: ar.NewArticleAdapter(),
	}
	return &interactor
}
