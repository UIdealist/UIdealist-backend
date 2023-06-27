package models

type User struct {
	Username string `db:"usrusername" json:"username" validate:"required,username,lte=255"`
	Password string `db:"usrpassword" json:"password" validate:"required,lte=255"`
}
