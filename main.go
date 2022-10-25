package main

import (
	"log"
	"rent-book/controllers"
	"rent-book/databases"
	"rent-book/repositories"
	"rent-book/routes"
	"rent-book/services"

	"github.com/labstack/echo/v4"
)

func main() {
	db := databases.GetConnectMysql()
	defer db.Close()

	//users
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	//books
	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookController := controllers.NewBookController(bookService)

	// rent book
	rentRepository := repositories.NewRentRepository(db)
	rentService := services.NewRentService(rentRepository)
	rentController := controllers.NewRentController(rentService)

	//login auth
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	e := echo.New()
	// user
	routes.UserPath(e, userController)
	// login
	routes.LoginAuth(e, authController)
	// book
	routes.BookPath(e, bookController)
	// rent
	routes.RentBook(e, rentController)

	//start server
	log.Fatal(e.Start(":9090"))
}
