package model

import "gorm.io/gorm"

type Layer struct {
	gorm.Model
	ID    string `gorm:"default:generated();" json:"id"`
	MapID string
	name  string
	tags  string
}
