package models

type LoginRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginResponse struct {
	UserId int    `json:"userid" db:"userid"`
	Name   string `json:"name" db:"name"`
	Token  string `json:"token" db:"token"`
}

type LoginDataResponse struct {
	UserId   int    `json:"userid" db:"userid"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
