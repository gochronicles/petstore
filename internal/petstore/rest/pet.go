package rest

import (
	"fmt"
	"net/http"
	db "petstore/internal/petstore/repo/postgres"
	"petstore/internal/petstore/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

//CreatePet route for POST
func CreatePet(c echo.Context) error {
	fmt.Println("Creating Pet")
	p := new(db.Pet)
	// bind request body to the model object
	if err := c.Bind(p); err != nil {
		panic(err)
	}
	err:=service.CreatePet(p)
	if err!=nil{
		panic(err)
	}
	return c.String(http.StatusCreated, "Pet created successfully")
}

//GetPetByCategory route for GET
func GetPetByCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("category_id"))
	response, err := service.GetPetByCategory(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, response)

}

//GetPet route for GET
func GetPet(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := service.GetPet(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, response)
}

//DeletePet route for DELETE
func DeletePet(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.DeletePet(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, "Deleted")

}
