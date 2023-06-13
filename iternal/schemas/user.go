package schemas

import "github.com/google/uuid"

type (
	UserInfo struct {
		ID       uuid.UUID `json:"id"  validate:"required"`
		Email    string    `json:"email"  validate:"required,email"`
		Username string    `json:"username"  validate:"required"`
	}

	UserUpdate struct {
		Email    string `json:"email"  validate:"required,email"`
		Username string `json:"username"  validate:"required"`
	}
)
