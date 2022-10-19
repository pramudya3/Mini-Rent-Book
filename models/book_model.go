package models

type Book struct {
	BookId int    `json:"bookid" db:"bookid"`
	Name   string `json:"name" db:"name"`
	Author string `json:"author" db:"author"`
}

type NewBook struct {
	Name   string `json:"name" db:"name"`
	Author string `json:"author" db:"author"`
}
