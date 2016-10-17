package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// RootMountHandler sets the routes for /
func RootMountHandler(base *echo.Group) {
	base.GET("", root)
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "This is the root!")
}
