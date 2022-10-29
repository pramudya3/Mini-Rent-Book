package models

type Book struct {
	BookId int    `json:"bookId" db:"bookId"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
	UserId int    `json:"userId" db:"userId"`
}

type NewBook struct {
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}

type BookResponse struct {
	BookId int    `json:"bookId" db:"bookId"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}
