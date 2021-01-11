package rest

import (
	"fmt"
	"net/http"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/internal/petstore/service"

	"github.com/labstack/echo/v4"
)

//CreatePet route for POST
func CreatePet(c echo.Context) error {
	fmt.Println("Creating Breed")
	b := db.Pet{ID: 1, Name: "abc"}
	service.CreatePet(&b)
	return c.String(http.StatusCreated, "Categories created successfully")
}
