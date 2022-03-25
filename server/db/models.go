package db

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type MapChunk struct {
	gorm.Model
	ID   string `gorm:"type:varchar(16);"`
	Name string `gorm:"unique"`
}

type User struct {
	gorm.Model
	ID   string `gorm:"type:varchar(16);"`
	Name string `gorm:"unique"`
}
