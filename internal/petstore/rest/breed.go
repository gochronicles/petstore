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
		panic(err)
	}
	// call the service
	err := service.CreateBreed(b)
	if err != nil {
		log.Println(err)
	}
	//return success response
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
