package models

import "time"

type User struct {
	UserId      int        `json:"userid" db:"userid"`
	Name        string     `json:"name" db:"name"`
	Email       string     `json:"email" db:"email"`
	Password    string     `json:"password" db:"password"`
	PhoneNumber string     `json:"phone_number" db:"phone_number"`
	Address     string     `json:"address" db:"address"`
	CreatedAt   time.Time  `json:"createdAt" db:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updatedAt"`
}

type NewUserRequest struct {
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Address     string    `json:"address" db:"address"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type UserResponse struct {
	UserId      int        `json:"userid" db:"userid"`
	Name        string     `json:"name" db:"name"`
	Email       string     `json:"email" db:"email"`
	PhoneNumber string     `json:"phone_number" db:"phone_number"`
	Address     string     `json:"address" db:"address"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updatedAt"`
}

type UpdateRequest struct {
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Address     string    `json:"address" db:"address"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updatedAt"`
}

type Login struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
