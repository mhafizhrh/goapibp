package model

type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
