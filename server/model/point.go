package model

import "gorm.io/gorm"

type Point struct {
	x float32
	y float32
}

type Location struct {
	gorm.Model
	ID    string `gorm:"default:generated();" json:"id"`
	point Point
	alias string
}
