package domain

type Article struct {
	ID       string `json:"id"`
	name     string
	text     string
	textHTML string
	tags     string
}
