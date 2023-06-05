package schemas

type (
	UserSignUp struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	UserSignIn struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	AccessToken struct {
		Token string `json:"access_token"`
	}
)
