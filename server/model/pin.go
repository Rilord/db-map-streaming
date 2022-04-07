package model

import (
	"gorm.io/gorm"
)

type PinType string

const (
	PinTypeStarredPlace   PinType = "STARRED"
	PinTypeFavouritePlace PinType = "FAVOURITE"
	PintTypeOwnList       PinType = "OWNLIST"
)

type Pin struct {
	gorm.Model
	ID         string `json:"id" gorm:"default:generated();"`
	pinType    PinType
	locationID string
	tags       string
	authorID   string
	postID     string
}
