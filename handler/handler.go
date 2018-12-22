package handler

import (
	"fmt"
	"net/http"

	"github.com/boomstarternetwork/bestore"
	"github.com/labstack/echo"
)

type Handler struct {
	store bestore.Store
}

func NewHandler(s bestore.Store) Handler {
	return Handler{
		store: s,
	}
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
