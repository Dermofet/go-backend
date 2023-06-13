package schemas

import (
	_ "github.com/go-playground/validator"
)

type (
	UserSignUp struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
		Username string `json:"username" validate:"required"`
	}

	UserSignIn struct {
		Password string `json:"password" validate:"required,min=6"`
		Email    string `json:"email" validate:"required,email"`
	}

	AccessToken struct {
		Token string `json:"access_token"`
	}
)
