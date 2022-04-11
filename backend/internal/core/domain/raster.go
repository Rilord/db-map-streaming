package domain

type Raster struct {
	ID           string `json:"id"`
	pathToRaster string
	name         string
	tileID       string
}
