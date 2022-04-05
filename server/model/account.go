package model

import "gorm.io/gorm"

type AccountType string

const (
	AccountTypeUser        AccountType = "USER"
	AccountTypeContributor AccountType = "CONTRIBUTOR"
	AccountTypeModerator   AccountType = "MODERATOR"
)

type Account struct {
	gorm.Model
	ID     string `gorm:"default:generated();" json:"id"`
	Type   AccountType
	UserId string
	User   User
	Bans   string
}
