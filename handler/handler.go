package handler

import (
	"fmt"
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

func (h Handler) ErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			err = fmt.Errorf("%v, %v", err, he.Internal)
		}
	} else {
		msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = echo.Map{"error": msg}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
	}
}
