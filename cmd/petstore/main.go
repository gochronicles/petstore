package main

import (
	"github.com/labstack/echo/v4"
	"petstore/internal/petstore/rest"
)

var (
	router = echo.New()
)

func main() {
	mapUrls()
	router.Logger.Fatal(router.Start(":1323"))
}

func mapUrls() {
	router.POST("/category",rest.CategoryEndpoint.CreateCategory)
}


