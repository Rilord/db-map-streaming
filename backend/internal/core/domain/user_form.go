package domain

import (
	"database/sql"
	"time"
)

type UserFormState string

const (
	UserFormStateCreated         UserFormState = "CREATED"
	UserFormStateOnConsideration UserFormState = "MODERATING"
	UserFormStateApproved        UserFormState = "APPROVED"
	UserFormStateDeclined        UserFormState = "DECLINED"
)

type UserForm struct {
	ID         string `json:"id"`
	UserID     string
	Name       *string `json:"name"`
	Password   *string `json:"phone"`
	Email      *string `json:"email"`
	SignedUpAt time.Time
	UpdatedAt  time.Time
	RemovedAt  sql.NullTime
}

type UserFormFilled struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

type UserFormFilter struct {
	State  []UserFormState `json:"state"`
	ID     []string        `json:"id"`
	UserID []string        `json:"userId"`
}
