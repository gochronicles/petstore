package main

import (
	"petstore/internal/petstore/rest"

	"github.com/labstack/echo/v4"
)

var (
	router = echo.New()
)

func main() {
	mapUrls()
	router.Logger.Fatal(router.Start(":3000"))
}

func mapUrls() {
	router.GET("/category", rest.GetAllCategory)
	router.GET("/category/:id", rest.GetCategory)
	router.POST("/breed", rest.CreateBreed)
	router.GET("/breed/:category_id", rest.GetBreedByCategory)
}
