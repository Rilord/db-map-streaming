package model

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	ID       string `gorm:"default:generated();" json:"id"`
	VectorID string
	name     string
	tags     string
}
