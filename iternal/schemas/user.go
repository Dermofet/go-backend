package schemas

import "github.com/google/uuid"

type (
	UserInfo struct {
		ID       uuid.UUID `json:"id"`
		Email    string    `json:"email"`
		Username string    `json:"username"`
	}

	UserUpdate struct {
		ID       uuid.UUID `json:"id"`
		Email    string    `json:"email"`
		Username string    `json:"username"`
	}
)
