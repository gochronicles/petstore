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
		panic(err)
	}
	// call category service
	id, err := service.CreateCategory(cat)
	if err != nil {
		panic(err)
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
	if err != nil {
		panic(err)
	}
	// send JSON response
	return c.JSON(http.StatusOK, response)

}

//GetCategory get a single category
func GetCategory(c echo.Context) error {
	// Get id from the path param for category. eg /category/1
	id, _ := strconv.Atoi(c.Param("id"))
	// Call category service to get category
	response, err := service.GetCategory(id)
	if err != nil {
		panic(err)
	}
	// send JSON response
	return c.JSON(http.StatusOK, response)

}
