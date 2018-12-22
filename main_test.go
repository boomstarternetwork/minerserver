package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boomstarternetwork/bestore"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const (
	runMode  = "testing"
	logLevel = "off"
)

func initTestingEntities() (*echo.Echo, *bestore.MockStore, error) {
	s := bestore.NewMockStore()
	e, err := initWebServer(s, runMode, logLevel)
	return e, s, err
}

func TestHandler_ProjectsList_success(t *testing.T) {
	e, s, err := initTestingEntities()
	if !assert.NoError(t, err) {
		return
	}

	s.On("GetProjects").Return([]bestore.Project{
		{ID: 1, Name: "name-1"},
		{ID: 2, Name: "name-2"},
	}, nil)

	const wantJSON = `{"projects":[{"id":1,"name":"name-1"},` +
		`{"id":2,"name":"name-2"}]}`

	req := httptest.NewRequest(http.MethodGet, "/projects", nil)
	res := httptest.NewRecorder()

	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, wantJSON, res.Body.String())
}
