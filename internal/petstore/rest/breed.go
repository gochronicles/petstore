package rest

import (
	"fmt"
	"net/http"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/internal/petstore/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

//CreateBreed route for POST
func CreateBreed(c echo.Context) error {
	fmt.Println("Creating Breed")
	b := new(db.Breed)
	if err := c.Bind(b); err != nil {
		panic(err)
	}
	err := service.CreateBreed(b)
	if err != nil {
		fmt.Println(err)
	}
	return c.String(http.StatusCreated, "Breed created successfully")
}

//GetBreedByCategory send category id
func GetBreedByCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("category_id"))
	response, err := service.GetBreedByCategory(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, response)
}
