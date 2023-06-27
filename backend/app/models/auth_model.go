package models

// SignIn struct to describe login user.
type SignIn struct {
	Username string `json:"username" validate:"required,username,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

// SignUp struct to describe register user.
type SignUp struct {
	Username string `json:"username" validate:"required,username,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`

	// We won't let anyone sign up without a secret password.
	SecretPassword string `json:"secretPassword" validate:"required,lte=255"`
}
