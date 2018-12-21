package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

func (h Handler) Projects(c echo.Context) error {
	projects, err := h.store.GetProjects()
	if err != nil {
		return errors.New("failed to list projects from store: " + err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"projects": projects,
	})
}
