package model

type Tile struct {
	ID        string `json:"id"`
	layerID   string
	bounds    VectorObject `json:"bounds"`
	zoomLevel VectorObject `json:"zoom"`
}
