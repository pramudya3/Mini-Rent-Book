package models

type Book struct {
	BookId        int    `json:"bookId" db:"bookId"`
	AddedByUser   int    `json:"addedByUser" db:"addedByUser"`
	Title         string `json:"title" db:"title"`
	Author        string `json:"author" db:"author"`
	UpdatedByUser int    `json:"updatedByUser" db:"updatedByUser"`
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

// type UpdateBookResponse struct {
// 	BookId       int    `json:"bookId" db:"bookId"`
// 	Title        string `json:"title" db:"title"`
// 	Author       string `json:"author" db:"author"`
// 	UpdateByUser int    `json:"updateByUser" db:"updateByUser"`
// }
