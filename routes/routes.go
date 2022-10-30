package routes

import (
	"rent-book/controllers"
	"rent-book/middlewares"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uc *controllers.UserController) {
	e.POST("/users", uc.CreateUser)
	e.GET("/users/:userId", uc.GetUserById)
	e.GET("/users", uc.GetAllUser)
	e.DELETE("/users", uc.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/users", uc.UpdateUser, middlewares.JWTMiddleware())

}

func BookPath(e *echo.Echo, uc *controllers.BookController) {
	e.POST("/books", uc.NewBook, middlewares.JWTMiddleware())
	e.GET("/books/:bookId", uc.GetBookById)
	e.GET("/books", uc.GetAllBook)
	e.DELETE("/books/:bookId", uc.DeleteBook, middlewares.JWTMiddleware())
	e.PUT("/books/:bookId", uc.UpdateBook, middlewares.JWTMiddleware())
}

func LoginAuth(e *echo.Echo, ac *controllers.AuthController) {
	e.POST("/login", ac.Login)
}

func RentBook(e *echo.Echo, rc *controllers.RentController) {
	e.POST("/rents", rc.NewRent, middlewares.JWTMiddleware())
	e.GET("/rents", rc.GetAllRent)
	e.PUT("/rents/:rentId", rc.UpdateRent, middlewares.JWTMiddleware())
}
