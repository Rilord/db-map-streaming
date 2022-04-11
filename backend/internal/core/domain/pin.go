package domain

type PinType string

const (
	PinTypeStarredPlace   PinType = "STARRED"
	PinTypeFavouritePlace PinType = "FAVOURITE"
	PintTypeOwnList       PinType = "USERS"
)

type Pin struct {
	ID         string `json:"id"`
	pinType    PinType
	locationID string
	tags       string
	authorID   string
	postID     string
}
