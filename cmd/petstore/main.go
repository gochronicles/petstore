package main

import (
	"net/http"
	"petstore/internal/petstore/rest"

	"github.com/labstack/echo/v4"
)

var (
	router = echo.New()
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to our Petstore App!!")
}

func main() {
	mapUrls()
	router.Logger.Fatal(router.Start(":3000"))
}

func mapUrls() {
	router.GET("/", ping)
	router.POST("/category", rest.CreateCategory)
	router.GET("/category", rest.GetAllCategory)
	router.GET("/category/:id", rest.GetCategory)
	router.POST("/breed", rest.CreateBreed)
	router.GET("/breed/:category_id", rest.GetBreedByCategory)
	router.POST("/location", rest.CreateLocation)
	router.GET("/location", rest.GetAllLocation)
	router.GET("/pet", rest.GetPetByCategory)
	router.POST("/pet", rest.CreatePet)
	router.GET("/pet/:id", rest.GetPet)
	router.DELETE("/pet", rest.DeletePet)
	router.POST("/pet/upload", rest.UploadPetImage)
}
