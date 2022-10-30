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

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var newUser models.NewUserRequest
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(err.Error()))
	}

	ctx := c.Request().Context()
	errCreateUser := uc.userService.CreateUser(ctx, newUser)
	if errCreateUser != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(errCreateUser.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData("success create user"))
}

func (uc *UserController) GetUserById(c echo.Context) error {

	idString := c.Param("userId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed("id not recognize"))
	}

	ctx := c.Request().Context()
	user, err := uc.userService.GetUserById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success get user by id", user))
}

func (uc *UserController) GetAllUser(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := uc.userService.GetAllUser(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success get all user", users))
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	// mendapatkan userId dari token
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("unauthorized"))
	}

	ctx := c.Request().Context()
	err := uc.userService.DeleteUser(ctx, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData("success delete user"))
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	// mendapatkan userId dari token
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("unauthorized"))
	}

	var updateUser models.UpdateRequest
	err := c.Bind(&updateUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(err.Error()))
	}

	ctx := c.Request().Context()
	user, err := uc.userService.UpdateUser(ctx, updateUser, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success update user", user))
}
