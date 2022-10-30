package models

import (
	"time"
)

type Rent struct {
	RentId     int       `json:"rentId" db:"rentId"`
	BookId     int       `json:"bookId" db:"bookId"`
	UserId     int       `json:"userId" db:"userId"`
	BorrowDate time.Time `json:"borrow_date" db:"borrow_date"`
	ReturnMax  time.Time `json:"return_max" db:"return_max"`
	ReturnDate time.Time `json:"return_date" db:"return_date"`
}

type NewRent struct {
	BookId     int       `json:"bookId" db:"bookId"`
	UserId     int       `json:"userId" db:"userId"`
	BorrowDate time.Time `json:"borrow_date" db:"borrow_date"`
	ReturnMax  time.Time `json:"return_max" db:"return_max"`
}

type NewRentResponse struct {
	BookId     int       `json:"bookId" db:"bookId"`
	UserId     int       `json:"userId" db:"userId"`
	BorrowDate time.Time `json:"borrow_date" db:"borrow_date"`
	ReturnMax  time.Time `json:"return_max" db:"return_max"`
}

type UpdateRent struct {
	ReturnDate time.Time `json:"return_date" db:"return_date"`
}
