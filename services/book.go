package services

import (
	"context"
	"errors"
	"rent-book/models"
	"rent-book/repositories"
)

type BookServiceInterface interface {
	NewBook(ctx context.Context, newBook models.NewBook, idToken int) error
	GetBookByIdLogin(ctx context.Context, bookId int, idToken int) (models.BookResponse, error)
	GetBookById(ctx context.Context, bookId int) (models.BookResponse, error)
	GetBookByTitle(ctx context.Context, title string) (models.BookResponse, error)
	GetAllBook(ctx context.Context) ([]models.Book, error)
	DeleteBook(ctx context.Context, idToken int) error
	UpdateBook(ctx context.Context, updateBook models.NewBook, bookId int, idToken int) (models.Book, error)
}

type BookService struct {
	bookRepository repositories.BookRepositoryInterface
}

func NewBookService(bookRepo repositories.BookRepositoryInterface) BookServiceInterface {
	return &BookService{
		bookRepository: bookRepo,
	}
}

func (bs *BookService) NewBook(ctx context.Context, newBook models.NewBook, idToken int) error {
	if newBook.Title == "" {
		return errors.New("Title Required")
	}
	if newBook.Author == "" {
		return errors.New("Author Required")
	}
	err := bs.bookRepository.NewBook(ctx, newBook, idToken)
	return err
}

func (bs *BookService) GetBookByIdLogin(ctx context.Context, bookId int, idToken int) (models.BookResponse, error) {
	book, err := bs.bookRepository.GetBookByIdLogin(ctx, bookId, idToken)

	bookResponse := models.BookResponse{
		BookId: book.BookId,
		Title:  book.Title,
		Author: book.Author,
	}
	return bookResponse, err
}

func (bs *BookService) GetBookById(ctx context.Context, bookId int) (models.BookResponse, error) {
	book, err := bs.bookRepository.GetBookById(ctx, bookId)

	bookResponse := models.BookResponse{
		BookId: book.BookId,
		Title:  book.Title,
		Author: book.Author,
	}
	return bookResponse, err
}

func (bs *BookService) GetBookByTitle(ctx context.Context, title string) (models.BookResponse, error) {
	book, err := bs.bookRepository.GetBookByTitle(ctx, title)

	bookResponse := models.BookResponse{
		BookId: book.BookId,
		Title:  book.Title,
		Author: book.Author,
	}
	return bookResponse, err
}

func (bs *BookService) GetAllBook(ctx context.Context) ([]models.Book, error) {
	books, err := bs.bookRepository.GetAllBook(ctx)
	return books, err
}

func (bs *BookService) DeleteBook(ctx context.Context, idToken int) error {
	err := bs.bookRepository.DeleteBook(ctx, idToken)
	return err
}

func (bs *BookService) UpdateBook(ctx context.Context, updateBook models.NewBook, bookId int, idToken int) (models.Book, error) {
	// getBook, err := bs.bookRepository.GetBookByIdLogin(ctx, bookId, idToken)
	getBook, err := bs.bookRepository.GetBookByIdLogin(ctx, bookId, idToken)
	if err != nil {
		return models.Book{}, err
	}
	if updateBook.Title != "" {
		getBook.Title = updateBook.Title
	}
	if updateBook.Author != "" {
		getBook.Author = updateBook.Author
	}

	book, err := bs.bookRepository.UpdateBook(ctx, getBook, bookId, idToken)

	responseUpdate := models.Book{
		BookId:        getBook.BookId,
		Title:         book.Title,
		Author:        book.Author,
		AddedByUser:   getBook.AddedByUser,
		UpdatedByUser: book.UpdatedByUser,
	}
	return responseUpdate, err
}
