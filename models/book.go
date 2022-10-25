package models

type Book struct {
	BookId int    `json:"bookId" db:"bookId"`
	UserId int    `json:"userId" db:"userId"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}

type NewBook struct {
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}
