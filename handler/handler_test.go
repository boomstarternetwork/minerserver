package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boomstarternetwork/minerserver/store"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHandler_ProjectsList_dbError(t *testing.T) {
	e := echo.New()

	ps := store.NewMockProjectsStore()
	ps.On("List").Return(nil, errors.New("test error"))

	const wantJSON = `{"error":"internal server error"}`

	api := NewHandler(ps)

	req := httptest.NewRequest(http.MethodGet, "/projects/list", nil)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if assert.NoError(t, api.ProjectsList(c)) {
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, wantJSON, res.Body.String())
	}
}

func TestHandler_ProjectsList_success(t *testing.T) {
	e := echo.New()

	ps := store.NewMockProjectsStore()
	ps.On("List").Return([]store.Project{
		{ID: "id-1", Name: "name-1"},
		{ID: "id-2", Name: "name-2"},
	}, nil)

	const wantJSON = `{"result":[{"id":"id-1","name":"name-1"},` +
		`{"id":"id-2","name":"name-2"}]}`

	api := NewHandler(ps)

	req := httptest.NewRequest(http.MethodGet, "/projects/list", nil)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if assert.NoError(t, api.ProjectsList(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, wantJSON, res.Body.String())
	}
}
