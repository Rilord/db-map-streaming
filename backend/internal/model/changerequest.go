package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ChangeRequest struct {
	gorm.Model
	ID     string         `gorm:"default:generated();" json:"id"`
	PinIDs pq.StringArray `gorm:"type:text[]"`
}
