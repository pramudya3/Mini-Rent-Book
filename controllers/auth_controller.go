package controllers

import (
	"net/http"
	"rent-book/helpers"
	"rent-book/models"
	"rent-book/services"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService services.AuthServiceInterface
}

func NewAuthController(authService services.AuthServiceInterface) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Login(c echo.Context) error {
	var userLogin models.LoginRequest
	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(err.Error()))
	}

	ctx := c.Request().Context()
	loginResponse, err := ac.authService.Login(ctx, userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccess(200, "succes login", loginResponse))
}
