package rest

import (
	"net/http"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/internal/petstore/service"

	"github.com/labstack/echo/v4"
)

//CreateLocation route for POST
func CreateLocation(c echo.Context) error {
	l := new(db.Location)
	if err := c.Bind(l); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "Error binding request body  ", err.Error())
	}
	err := service.CreateLocation(l)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error in creating a location : "+err.Error())
	}
	return c.String(http.StatusCreated, "Location created successfully")
}

//GetAllLocation get all categories
func GetAllLocation(c echo.Context) error {
	response, err := service.GetAllLocation()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error in getting locations : "+err.Error())

	}
	return c.JSON(http.StatusOK, response)

}
