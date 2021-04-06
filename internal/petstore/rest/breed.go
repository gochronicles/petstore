package rest

import (
	"log"
	"net/http"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/internal/petstore/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

//CreateBreed : POST route for creating a new Breeds
func CreateBreed(c echo.Context) error {
	log.Println("Creating Breed")
	// create a new object of breed model
	b := new(db.Breed)
	// bind request body to the model object
	if err := c.Bind(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error in Binding request object "+err.Error())
	}
	// call the service
	err := service.CreateBreed(b)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in Creating a breed object "+err.Error())
	}
	//return success response
	return c.String(http.StatusCreated, "Breed created successfully")
}

//GetBreedByCategory send category id
func GetBreedByCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id sent is incorrect. Please send a valid integer value")
	}
	response, err := service.GetBreedByCategory(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in getting a breed by category "+err.Error())

	}
	return c.JSON(http.StatusOK, response)
}
