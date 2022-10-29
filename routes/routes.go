package routes

import (
	"rent-book/controllers"
	"rent-book/middlewares"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uc *controllers.UserController) {
	e.POST("/user", uc.CreateUser)
	e.GET("/user/:userId", uc.GetUserById)
	e.GET("/users", uc.GetAllUser)
	e.DELETE("/user", uc.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/user", uc.UpdateUser, middlewares.JWTMiddleware())

}

func BookPath(e *echo.Echo, uc *controllers.BookController) {
	e.POST("/book", uc.NewBook, middlewares.JWTMiddleware())
	e.GET("/book/:bookId", uc.GetBookById)
	e.GET("/books/:title", uc.GetBookByTitle)
	e.GET("/books", uc.GetAllBook)
	e.DELETE("/book/:bookId", uc.DeleteBook, middlewares.JWTMiddleware())
	e.PUT("/book/:bookId", uc.UpdateBook, middlewares.JWTMiddleware())
}

func LoginAuth(e *echo.Echo, ac *controllers.AuthController) {
	e.POST("/login", ac.Login)
}

func RentBook(e *echo.Echo, rc *controllers.RentController) {
	e.POST("/rent", rc.NewRent, middlewares.JWTMiddleware())
	e.GET("/rent", rc.GetRentByLogin, middlewares.JWTMiddleware())
	e.GET("/rentAll", rc.GetAllRent)
	e.PUT("/rent", rc.UpdateRent, middlewares.JWTMiddleware())
}
