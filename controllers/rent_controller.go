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

type RentController struct {
	rentService services.RentServiceInterface
}

func NewRentController(rentService services.RentServiceInterface) *RentController {
	return &RentController{
		rentService: rentService,
	}
}

func (rc *RentController) NewRent(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("Unauthorized"))
	}

	var newRent models.NewRent
	err := c.Bind(&newRent)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(err.Error()))
	}
	ctx := c.Request().Context()
	errNewRent := rc.rentService.NewRent(ctx, newRent, idToken)
	if errNewRent != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(errNewRent.Error()))
	}
	return c.JSON(http.StatusOK, helpers.APIResponseSuccessWithoutData("success add new rent"))
}

func (rc *RentController) GetAllRent(c echo.Context) error {
	ctx := c.Request().Context()
	rents, err := rc.rentService.GetAllRent(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success get all rent", rents))

}

func (rc *RentController) UpdateRent(c echo.Context) error {
	idString := c.Param("rentId")
	rentId, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed("RentId not found"))
	}
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, helpers.APIResponseFailed("unauthorized"))
	}

	var updateRent models.UpdateRent
	errRent := c.Bind(&updateRent)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.APIResponseFailed(errRent.Error()))
	}

	ctx := c.Request().Context()
	rent, err := rc.rentService.UpdateRent(ctx, updateRent, rentId, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.APIResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.APIResponseSuccess("success update rent", rent))
}
