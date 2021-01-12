package rest

import (
	"fmt"
	"net/http"
	"petstore/internal/petstore/repo"
	"petstore/internal/petstore/service"

	"github.com/labstack/echo/v4"
)

//CreatePet route for POST
func CreatePet(c echo.Context) error {
	fmt.Println("Creating Breed")
	b := repo.Pet{ID: 1, Name: "abc"}
	service.CreatePet(&b)
	return c.String(http.StatusCreated, "Categories created successfully")
}
