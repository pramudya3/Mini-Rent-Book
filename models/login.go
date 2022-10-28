package models

type LoginRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginResponse struct {
	Token  string `json:"token" db:"token"`
	UserId int    `json:"userid" db:"userid"`
	Name   string `json:"name" db:"name"`
}

type LoginDataResponse struct {
	UserId   int    `json:"userid" db:"userid"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
