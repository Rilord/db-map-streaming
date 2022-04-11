package domain

type GeometryType int8

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
	ID       string `json:"id"`
	geomType GeometryType
	data     []float32
	tileID   string
}
