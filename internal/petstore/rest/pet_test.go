package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"

	db "petstore/internal/petstore/repo/postgres"
	"petstore/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	petData = db.Pet{
	Name: "some_name",
	Age:1,
	Description: "some description",
	PetLocation: models.Location{ID:1},
	PetCategory:models.Category{ID:1},
	PetBreed: models.Breed{ID:1},
}
	petJSON = `{"name":"some_name","age":1,"description":"some description"},"location":{"id":1},"category":{"id":1},"breed":{"id":1,"category_id":1}`
)

func TestCreatePet(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/pet", strings.NewReader(petJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)
	if assert.NoError(t, CreatePet(ctx)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(t, recorder.Body.String(), "Pet created successfully")

	}
}

func TestGetPetByCategory(t *testing.T) {
	e := echo.New()
	response := []db.Pet{}
	request := httptest.NewRequest(http.MethodGet, "/pet", nil)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)
	ctx.SetPath("/pet")
	ctx.SetParamNames("category_id")
	ctx.SetParamValues("1")
	fmt.Println(ctx.QueryParam("category_id"))
	// Assertions
	if assert.NoError(t, GetPetByCategory(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NotNil(t, recorder.Body.String())
		assert.NotEmpty(t, response)
	}
}