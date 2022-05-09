package comment

import "game-server/internal/core/domain"

func (i *Interactor) Add(comment domain.Comment) error {
	return i.Comment.Add(comment)
}
