package models

import "time"

type NewRent struct {
	BookId     int       `json:"bookId" db:"bookId"`
	BorrowDate time.Time `json:"borrow_date" db:"borrow_date"`
	ReturnDate time.Time `json:"return_date" db:"return_date"`
}

type UpdateRent struct {
	ReturnDate time.Time `json:"return_date" db:"return_date"`
}

type Rent struct {
	BookId     int       `json:"bookId" db:"bookId"`
	UserId     int       `json:"userId" db:"userId"`
	BorrowDate time.Time `json:"borrow_date" db:"borrow_date"`
	ReturnDate time.Time `json:"return_date" db:"return_date"`
}
