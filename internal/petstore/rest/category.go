package rest

import (
	"net/http"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/internal/petstore/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

//CreateCategory route for POST
func CreateCategory(c echo.Context) error {
	// create a new object of category model
	cat := new(db.Category)
	// bind request body to the model object
	if err := c.Bind(cat); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "Error binding request body  ", err.Error())

	}
	// call category service
	id, err := service.CreateCategory(cat)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in creating a category"+err.Error())
	}
	// create a response
	response := map[string]interface{}{"id": id}
	//return success response
	return c.JSON(http.StatusCreated, response)
}

//GetAllCategory get all categories
func GetAllCategory(c echo.Context) error {
	// Call category service to get all categories
	response, err := service.GetAllCategory()
	if len(response) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "No categories found")
	}
	if err != nil {
		panic(err)
	}
	// send JSON response
	return c.JSON(http.StatusOK, response)

}

//GetCategory get a single category
func GetCategory(c echo.Context) error {
	// Get id from the path param for category. eg /category/1
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id sent is incorrect. Please send a valid integer value")
	}
	// Call category service to get category
	response, err := service.GetCategory(id)
	if response == nil {
		return echo.NewHTTPError(http.StatusNotFound, "No records found for category id : "+strconv.Itoa(id))
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failure in fetching category")
	}
	// send JSON response
	return c.JSON(http.StatusOK, response)

}
