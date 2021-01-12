package rest

import (
	"fmt"
	"net/http"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/internal/petstore/service"

	"github.com/labstack/echo/v4"
)

//CreateBreed route for POST
func CreateBreed(c echo.Context) error {
	fmt.Println("Creating Breed")
	b := db.Breed{BreedName: "abc", CategoryID: 1}
	err := service.CreateBreed(&b)
	if err != nil {
		fmt.Println(err)
	}
	return c.String(http.StatusCreated, "Breed created successfully")
}
