package rest

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/internal/petstore/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

var imageDir string = "images"

//CreatePet route for POST
func CreatePet(c echo.Context) error {
	fmt.Println("Creating Pet")
	p := new(db.Pet)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error in Binding request object "+err.Error())
	}
	err := service.CreatePet(p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in creating a pet "+err.Error())
	}
	return c.String(http.StatusCreated, "Pet created successfully")
}

//GetPetByCategory route for GET
func GetPetByCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("category_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "category id sent is incorrect. Please send a valid integer value")
	}
	response, err := service.GetPetByCategory(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in creating a pet "+err.Error())

	}
	return c.JSON(http.StatusOK, response)

}

//GetPet route for GET
func GetPet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id sent is incorrect. Please send a valid integer value")
	}
	response, err := service.GetPet(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in fetching a pet by id"+err.Error())

	}
	return c.JSON(http.StatusOK, response)
}

//DeletePet route for DELETE
func DeletePet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id sent is incorrect. Please send a valid integer value")
	}
	err = service.DeletePet(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error in deleting a pet "+err.Error())

	}
	return c.JSON(http.StatusOK, "Deleted")

}

func UploadPetImage(c echo.Context) error {
	// Multipart form
	id, err := strconv.Atoi(c.FormValue("id"))
	// Update database
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't read file from form value")
	}
	// Source
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't open file")
	}
	defer src.Close()
	// Destination
	dstPath := filepath.Join(imageDir, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't create destination file")
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't copy file")
	}
	err = service.UpdateImageURL(&db.Pet{ID: id, ImageURL: dstPath})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't read file")
	}
	return c.JSON(http.StatusCreated, "File uploaded successfully")
}
