package model

import (
	"gorm.io/gorm"
)

type Tile struct {
	gorm.Model
	ID        string `gorm:"default:generated();" json:"id"`
	layerID   string
	bounds    GeometryObject
	zoomLevel GeometryObject
}
