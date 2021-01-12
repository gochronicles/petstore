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
	cat := new(db.Category)
	if err := c.Bind(cat); err != nil {
		panic(err)
	}
	err := service.CreateCategory(cat)
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusCreated, "Categories created successfully")
}

//GetAllCategory get all categories
func GetAllCategory(c echo.Context) error {
	response, err := service.GetAllCategory()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, response)

}

//GetCategory get a single category
func GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := service.GetCategory(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, response)

}
