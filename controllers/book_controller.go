package controllers

import (
	"net/http"
	"rent-book/helpers"
	"rent-book/middlewares"
	"rent-book/models"
	"rent-book/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookService services.BookServiceInterface
}

func NewBookController(bookService services.BookServiceInterface) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

func (bc *BookController) NewBook(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("unauthorized"))
	}

	var newBook models.NewBook
	err := c.Bind(&newBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(err.Error()))
	}
	ctx := c.Request().Context()
	errNewBook := bc.bookService.NewBook(ctx, newBook, idToken)
	if errNewBook != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(errNewBook.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData("succes add a new book"))
}

func (bc *BookController) GetBookById(c echo.Context) error {
	idString := c.Param("bookId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed("BookId not found"))
	}

	ctx := c.Request().Context()
	book, err := bc.bookService.GetBookById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success get book by id", book))
}

func (bc *BookController) GetAllBook(c echo.Context) error {
	ctx := c.Request().Context()
	books, err := bc.bookService.GetAllBook(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success get all books", books))
}

func (bc *BookController) DeleteBook(c echo.Context) error {
	idString := c.Param("bookId")
	bookId, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed("BookId not found"))
	}

	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("unauthorized"))
	}
	ctx := c.Request().Context()
	errDelete := bc.bookService.DeleteBook(ctx, bookId, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(errDelete.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData("success delete a book"))
}

func (bc *BookController) UpdateBook(c echo.Context) error {
	idString := c.Param("bookId")
	bookId, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed("BookId not found"))
	}

	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("unauthorized"))
	}

	var updateBook models.NewBook
	errUpdate := c.Bind(&updateBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(errUpdate.Error()))
	}

	ctx := c.Request().Context()
	book, err := bc.bookService.UpdateBook(ctx, updateBook, bookId, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success update book", book))
}
