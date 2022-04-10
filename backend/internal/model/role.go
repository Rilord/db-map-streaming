package model

type RoleType string

const (
	ModeratorRoleType   RoleType = "MODERATOR"
	ContributorRoleType RoleType = "CONTRIBUTOR"
	UserRoleType        RoleType = "USER"
)

type Role struct {
	Type   RoleType
	UserID string
}
