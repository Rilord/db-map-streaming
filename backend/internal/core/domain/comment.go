package domain

type Comment struct {
	ID        string `json:"id"`
	Text      string
	UserID    string
	ArticleID string
	ParentID  string
}
