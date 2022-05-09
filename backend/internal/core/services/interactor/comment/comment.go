package comment

import (
	ar "game-server/internal/adapter/repository"
	pr "game-server/internal/core/ports/repository"
)

type Interactor struct {
	Comment pr.CommentRepository
}

func NewInteractor() *Interactor {
	return &Interactor{
		Comment: ar.NewCommentAdapter(),
	}
}
