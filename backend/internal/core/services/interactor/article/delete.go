package article

func (i *Interactor) Delete(article_id string) error {
	return i.Article.Delete(article_id)
}
