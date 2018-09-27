package handler

import (
	"net/http"

	"bitbucket.org/boomstarternetwork/minerserver/store"
	"github.com/labstack/echo"
)

type Handler struct {
	projects store.ProjectsStore
}

func NewHandler(ps store.ProjectsStore) Handler {
	return Handler{
		projects: ps,
	}
}

func (h Handler) ProjectsList(c echo.Context) error {
	projects, err := h.projects.List()
	if err != nil {
		c.Logger().Error(err)
		c.JSON(http.StatusInternalServerError,
			echo.Map{"error": "internal server error"})
		return nil
	}

	c.JSON(http.StatusOK, echo.Map{
		"result": projects,
	})

	return nil
}
