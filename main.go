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

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	e := echo.New()
	routes.UserPath(e, userController)

	log.Fatal(e.Start(":9090"))
}
