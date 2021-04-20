package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	db "petstore/internal/petstore/repo/postgres"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	breedData = db.Breed{BreedName: "some breed", CategoryID: 1}
	breedJSON = `{"breedName":"some breed","categoryID":1}`
)

func TestCreateBreed(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/breed", strings.NewReader(breedJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)
	if assert.NoError(t, CreateBreed(ctx)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(t, recorder.Body.String(), "Breed created successfully")

	}
}

func TestGetBreedByCategory(t *testing.T) {
	e := echo.New()
	response := []db.Breed{}
	request := httptest.NewRequest(http.MethodGet, "/breed", nil)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)
	ctx.SetPath("/breed/:category_id")
	ctx.SetParamNames("category_id")
	ctx.SetParamValues("1")

	// Assertions
	if assert.NoError(t, GetBreedByCategory(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NotNil(t, recorder.Body.String())
		assert.NotEmpty(t, response)
	}
}
