package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	db "petstore/internal/petstore/repo/postgres"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	data         = db.Category{CategoryName: "rabbit"}
	categoryJSON = `{"categoryname":"rabbit"}`
)

func TestCreateCategory(t *testing.T) {
	response := db.Category{}
	e := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/category", strings.NewReader(categoryJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)
	if assert.NoError(t, CreateCategory(ctx)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
		json.Unmarshal(recorder.Body.Bytes(), &response)
		require.NotZero(t, response.ID)

	}
}

func TestGetCategory(t *testing.T) {
	e := echo.New()
	response := db.Category{}
	request := httptest.NewRequest(http.MethodGet, "/category", nil)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)
	ctx.SetPath("/category/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	// Assertions
	if assert.NoError(t, GetCategory(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NotNil(t, recorder.Body.String())
		assert.Equal(t, strconv.Itoa(response.ID), ctx.Param("id"))
		assert.NotNil(t, response.CategoryName)
	}
}
func TestGetAllCategory(t *testing.T) {
	e := echo.New()
	response := []db.Category{}
	request := httptest.NewRequest(http.MethodGet, "/category", nil)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)

	// Assertions
	if assert.NoError(t, GetAllCategory(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NotNil(t, recorder.Body.String())
		assert.NotEmpty(t, response)

	}

}
