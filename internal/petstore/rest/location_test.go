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
	locationData = db.Location{LocationName: "some_location"}
	locationJSON = `{"locationName":"some_location"}`
)

func TestCreateLocation(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/location", strings.NewReader(locationJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)
	if assert.NoError(t, CreateLocation(ctx)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(t, recorder.Body.String(), "Location created successfully")

	}
}
func TestGetAllLocation(t *testing.T) {
	e := echo.New()
	response := []db.Location{}
	request := httptest.NewRequest(http.MethodGet, "/location", nil)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)

	// Assertions
	if assert.NoError(t, GetAllLocation(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NotNil(t, recorder.Body.String())
		assert.NotEmpty(t, response)

	}

}
