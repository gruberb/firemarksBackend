package handlers

import (
	"firemarksBackend/models"
	"net/http"

	"github.com/labstack/echo"
)

// UsersMountHandler sets the routes for the User resource
func UsersMountHandler(base *echo.Group) {
	base.GET("", listUsers)
}

// GET /users
func listUsers(c echo.Context) error {
	results := &[]models.User{}

	// Find Users
	if err := models.QueryUsers(results); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, results)
}
