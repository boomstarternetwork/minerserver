package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boomstarternetwork/minerserver/store"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const (
	runMode  = "testing"
	logLevel = "off"
)

func initTestingEntities() (*echo.Echo, *mockStore, error) {
	s := newMockStore()
	e, err := initWebServer(s, runMode, logLevel)
	return e, s, err
}

func TestHandler_ProjectsList_success(t *testing.T) {
	e, s, err := initTestingEntities()
	if !assert.NoError(t, err) {
		return
	}

	s.On("ListProjects").Return([]store.Project{
		{ID: "id-1", Name: "name-1"},
		{ID: "id-2", Name: "name-2"},
	}, nil)

	const wantJSON = `{"projects":[{"id":"id-1","name":"name-1"},` +
		`{"id":"id-2","name":"name-2"}]}`

	req := httptest.NewRequest(http.MethodGet, "/projects", nil)
	res := httptest.NewRecorder()

	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, wantJSON, res.Body.String())
}
