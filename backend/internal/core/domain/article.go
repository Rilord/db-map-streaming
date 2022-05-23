package domain

type Article struct {
	ID       *string `json:"id"`
	name     *string
	text     *string
	textHTML *string
	tags     *string
}

func NewArticle(id *string, name *string, text *string, textHTML *string, tags *string) *Article {
	return &Article{ID: id, name: name, text: text, textHTML: textHTML, tags: tags}
}
