package model

import "gorm.io/gorm"

type Map struct {
	gorm.Model
	ID   string `gorm:"default:generated();" json:"id"`
	name string
	tags string
}
