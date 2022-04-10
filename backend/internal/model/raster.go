package model

import "gorm.io/gorm"

type Raster struct {
	gorm.Model
	ID           string `gorm:default:generated();"json:"id"`
	pathToRaster string
	name         string
	tileID       string
}
