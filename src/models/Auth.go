package models

type SignIn struct {
	Username string `json:"username" validate:"required,min=3,max=12"`
	Password string `json:"password" validate:"required,min=3,max=12"`
}

type TokenModel struct {
	Token string `json:"token"`
}
