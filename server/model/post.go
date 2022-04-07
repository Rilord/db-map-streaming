package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ID       string `gorm:default:generated();"json:"id"`
	name     string
	text     string
	htmlText string
	tags     string
}
