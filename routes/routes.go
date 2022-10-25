package routes

import (
	"rent-book/controllers"
	"rent-book/middlewares"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uc *controllers.UserController) {
	//CRUD users
	e.POST("/user", uc.CreateUser)
	e.GET("/user/:user_id", uc.GetUserById)
	e.GET("/users", uc.GetAllUser)
	e.DELETE("/user", uc.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/user", uc.UpdateUser, middlewares.JWTMiddleware())

}

func BookPath(e *echo.Echo, uc *controllers.BookController) {
	//CRUD books
	e.POST("/book", uc.NewBook, middlewares.JWTMiddleware())
	e.GET("book/:book_id", uc.GetBookById)
	e.GET("books", uc.GetAllBook)
	e.DELETE("book", uc.DeleteBook, middlewares.JWTMiddleware())
	e.PUT("/book", uc.UpdateBook, middlewares.JWTMiddleware())
}

func LoginAuth(e *echo.Echo, ac *controllers.AuthController) {
	e.POST("/login", ac.Login)
}

func RentBook(e *echo.Echo, rc *controllers.RentController) {
	e.POST("/rent", rc.NewRent, middlewares.JWTMiddleware())
	e.GET("/rent/:id", rc.GetRentById)
	e.GET("/rents", rc.GetAllRent)
	e.PUT("/rent", rc.UpdateRent, middlewares.JWTMiddleware())
}
