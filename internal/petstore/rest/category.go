package rest

import (
	"fmt"
	"net/http"
	"petstore/internal/petstore/repo"
	"petstore/internal/petstore/service"

	"github.com/labstack/echo/v4"
)

//CreateCategory route for POST
func CreateCategory(c echo.Context) error {
	fmt.Println("Creating Breed")
	b := repo.Category{ID: 1, CategoryName: "abc"}
	service.CreateCategory(&b)
	return c.String(http.StatusCreated, "Categories created successfully")
}
