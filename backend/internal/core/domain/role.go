package domain

type RoleType string

const (
	ModeratorRoleType   RoleType = "MODERATOR"
	ContributorRoleType RoleType = "CONTRIBUTOR"
	UserRoleType        RoleType = "USER"
)

type Role struct {
	ID     string `json:"id"`
	Type   RoleType
	UserID string
}
