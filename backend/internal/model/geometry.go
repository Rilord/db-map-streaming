package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type GeometryType int64

const (
	POINT           GeometryType = 0
	LINESTRING      GeometryType = 1
	POLYGON         GeometryType = 2
	POLYGONWITHHOLE GeometryType = 3
	COLLECTION      GeometryType = 4
	TRIANGLE        GeometryType = 5
	MULTIPOLYGON    GeometryType = 6
)

type VectorObject struct {
	gorm.Model
	ID string `gorm:default:generated();"json:"id"`

	geomType GeometryType
	data     pq.Float32Array `gorm:"type:float[]"`
	tileID   string
}
