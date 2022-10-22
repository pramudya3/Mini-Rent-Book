package routes

import (
	"rent-book/controllers"
	"rent-book/middlewares"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uc *controllers.UserController) {
	e.POST("/user", uc.CreateUser)
	e.GET("/user/:user_id", uc.GetUserById)
	e.GET("/users", uc.GetAllUser)
	e.DELETE("/user", uc.DeleteUser, middlewares.JWTMiddleware())
	// e.PUT("/user", uc.UpdateUser, middlewares.JWTMiddleware())
	e.PUT("/user/:user_id", uc.UpdateUser)

}
