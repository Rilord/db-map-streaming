package model

import (
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

type GeometryObject struct {
	geomType GeometryType
	data     []float32
}

type VectorObject struct {
	gorm.Model
	ID       string `gorm:default:generated();"json:"id"`
	geometry GeometryObject
	tileID   string
}

func (vec *VectorObject) VectorGeometryType() string {
	if vec == nil {
		return "UNDEFINED"
	}

	switch vec.geometry.geomType {
	case 0:
		return "POINT"
	case 1:
		return "LINESTRING"
	case 2:
		return "POLYGON"
	case 3:
		return "POLYGONWITHHOLE"
	case 4:
		return "COLLECTION"
	case 5:
		return "TRIANGLE"
	case 6:
		return "MULTIPOLYGON"
	default:
		return "UNDEFINED"

	}
}
