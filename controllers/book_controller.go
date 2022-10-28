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
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData("succes add new book"))
}

func (bc *BookController) GetBookByIdLogin(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("unauthorized"))
	}

	idString := c.Param("bookId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed("BookId not found"))
	}

	ctx := c.Request().Context()
	book, err := bc.bookService.GetBookByIdLogin(ctx, id, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success get book", book))
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
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success get book", book))
}

func (bc *BookController) GetBookByTitle(c echo.Context) error {
	title := c.Param("title")

	ctx := c.Request().Context()
	book, err := bc.bookService.GetBookByTitle(ctx, title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success get book by title", book))
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
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("unauthorized"))
	}
	ctx := c.Request().Context()
	err := bc.bookService.DeleteBook(ctx, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData("success delete book"))
}

func (bc *BookController) UpdateBook(c echo.Context) error {
	idString := c.Param("bookId")
	id, err := strconv.Atoi(idString)
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
	book, err := bc.bookService.UpdateBook(ctx, updateBook, id, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success update book", book))
}
