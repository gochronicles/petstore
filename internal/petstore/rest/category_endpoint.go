package rest

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"petstore/internal/petstore/repo"
	"petstore/internal/petstore/service"
)

var (
	CategoryEndpoint = categoryEndpoint{}
)

type categoryEndpoint struct{}

func (categoryEndpoint categoryEndpoint) CreateCategory(c echo.Context) error {
	// Get name and email

	requestCategory := new(repo.Category)
	json_map := make(map[string]interface{})
	error := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if error != nil {
		return error
	} else {
	}

	err := c.Bind(requestCategory)
	if err != nil {
		return err
	}

	if err = service.CategoryService.Save(requestCategory); err != nil {
		return err
	}
	return c.String(http.StatusCreated, "Categories created successfully")
}