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
	b := db.Breed{ID: 1, BreedName: "abc"}
	service.CreateBreed(&b)
	return c.String(http.StatusCreated, "Categories created successfully")
}
