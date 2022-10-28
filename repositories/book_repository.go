package repositories

import (
	"context"
	"database/sql"
	"errors"
	"rent-book/models"
)

type BookRepositoryInterface interface {
	NewBook(ctx context.Context, newBook models.NewBook, idToken int) error
	GetBookByIdLogin(ctx context.Context, bookId int, idToken int) (models.Book, error)
	GetBookById(ctx context.Context, bookId int) (models.Book, error)
	GetBookByTitle(ctx context.Context, title string) (models.Book, error)
	GetAllBook(ctx context.Context) ([]models.Book, error)
	DeleteBook(ctx context.Context, idToken int) error
	UpdateBook(ctx context.Context, updateBook models.Book, bookId int, idToken int) (models.Book, error)
}

type BookRepository struct {
	mysql *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		mysql: db,
	}
}

func (br *BookRepository) NewBook(ctx context.Context, newBook models.NewBook, idToken int) error {
	query := "INSERT INTO books (title, author, addedByUser) VALUES (?, ?, ?)"

	_, err := br.mysql.ExecContext(ctx, query, newBook.Title, newBook.Author, idToken)
	if err != nil {
		return err
	}

	return nil
}

func (br *BookRepository) GetBookByIdLogin(ctx context.Context, bookId int, idToken int) (models.Book, error) {
	var book models.Book
	query := "SELECT bookId, title, author FROM books WHERE bookId = ?"

	err := br.mysql.QueryRowContext(ctx, query, bookId).Scan(&book.BookId, &book.Title, &book.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, err
		}
		return models.Book{}, err
	}
	return book, nil
}

func (br *BookRepository) GetBookById(ctx context.Context, bookId int) (models.Book, error) {
	var book models.Book
	query := "SELECT bookId, title, author FROM books WHERE bookId = ?"

	err := br.mysql.QueryRowContext(ctx, query, bookId).Scan(&book.BookId, &book.Title, &book.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, err
		}
		return models.Book{}, err
	}
	return book, nil
}

func (br *BookRepository) GetBookByTitle(ctx context.Context, title string) (models.Book, error) {
	var book models.Book
	query := "SELECT bookId, title, author FROM books WHERE title = ?"

	err := br.mysql.QueryRowContext(ctx, query, title).Scan(&book.BookId, &book.Title, &book.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, err
		}
		return models.Book{}, err
	}
	return book, nil
}

func (br *BookRepository) GetAllBook(ctx context.Context) ([]models.Book, error) {
	query := "SELECT bookId, title, author, addedByUser FROM books"

	rows, err := br.mysql.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.AddedByUser)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (br *BookRepository) DeleteBook(ctx context.Context, idToken int) error {
	query := "DELETE FROM books WHERE bookId = ?"

	result, err := br.mysql.ExecContext(ctx, query, idToken)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("Book Not Found")
	}
	return nil
}

func (br *BookRepository) UpdateBook(ctx context.Context, updateBook models.Book, bookId int, idToken int) (models.Book, error) {
	query := "UPDATE books SET title = ?, author = ?, updatedByUser = ? WHERE bookId = ?"

	result, err := br.mysql.ExecContext(ctx, query, updateBook.Title, updateBook.Author, idToken, bookId)
	if err != nil {
		return models.Book{}, err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return models.Book{}, errors.New("Book Not Found")
	}
	return updateBook, nil
}
