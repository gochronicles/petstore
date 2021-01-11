package rest

import (
	"fmt"
	"net/http"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/internal/petstore/service"

	"github.com/labstack/echo/v4"
)

//CreateCategory route for POST
func CreateCategory(c echo.Context) error {
	fmt.Println("Creating Breed")
	b := db.Category{ID: 1, CategoryName: "abc"}
	service.CreateCategory(&b)
	return c.String(http.StatusCreated, "Categories created successfully")
}
