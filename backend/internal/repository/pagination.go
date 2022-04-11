package repository

import (
	"gorm.io/gorm"
)

func paginationQueryByCreatedAtDesc(query *gorm.DB, first *int, after *string) (*gorm.DB, error) {
	if after != nil {

	}

}
