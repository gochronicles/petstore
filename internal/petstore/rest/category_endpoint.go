package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	CategoryEndpoint = categoryEndpoint{}
)

type categoryEndpoint struct{}

func (categoryEndpoint categoryEndpoint) CreateCategory(c echo.Context) error {
	// Get name and email
	fmt.Println("Inside Create Category")
	return c.String(http.StatusCreated, "Categories created successfully")
}